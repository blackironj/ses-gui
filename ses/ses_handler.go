package ses

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	sessdk "github.com/aws/aws-sdk-go/service/ses"
)

const (
	charSet = "UTF-8"
)

// ListSEStemplates gets email-templates from AWS-SES
func ListSEStemplates() ([]*sessdk.TemplateMetadata, error) {
	if AwsSession == nil {
		return nil, errors.New("fail to access")
	}
	sesClient := sessdk.New(AwsSession)

	input := sessdk.ListTemplatesInput{
		MaxItems: aws.Int64(10),
	}

	templateOutputs := make([]*sessdk.TemplateMetadata, 0, 20)
	for {
		output, err := sesClient.ListTemplates(&input)
		if err != nil {
			return nil, err
		}

		templateOutputs = append(templateOutputs, output.TemplatesMetadata...)
		if output.NextToken == nil {
			break
		}
		input.NextToken = output.NextToken
	}

	return templateOutputs, nil
}

// UploadSEStemplate uploads a email-template to AWS-SES
func UploadSEStemplate(sesTemplate *sessdk.Template) error {
	sesClient := sessdk.New(AwsSession)

	createTemplateInput := &ses.CreateTemplateInput{
		Template: sesTemplate,
	}

	_, err := sesClient.CreateTemplate(createTemplateInput)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSEStemplate deletes a email-template from AWS-SES
func DeleteSEStemplate(name *string) error {
	sesClient := sessdk.New(AwsSession)

	deleteTemplateInput := &ses.DeleteTemplateInput{
		TemplateName: name,
	}

	_, err := sesClient.DeleteTemplate(deleteTemplateInput)
	if err != nil {
		return err
	}
	return nil
}

// GetSEStemplate gets a specific email-template from AWS-SES
func GetSEStemplate(name *string) (*sessdk.GetTemplateOutput, error) {
	sesClient := sessdk.New(AwsSession)
	return getSEStemplate(sesClient, name)
}

func getSEStemplate(sesClient *sessdk.SES, name *string) (*sessdk.GetTemplateOutput, error) {
	getTemplateInput := &ses.GetTemplateInput{
		TemplateName: name,
	}

	getTemplateOutput, err := sesClient.GetTemplate(getTemplateInput)
	if err != nil {
		return nil, err
	}
	return getTemplateOutput, nil
}

// SendEmailWithTemplate sends a email with html template
func SendEmailWithTemplate(sender, recipient, templateName string, datas ...map[string]interface{}) error {
	sesClient := sessdk.New(AwsSession)

	dest := &ses.Destination{
		CcAddresses: []*string{},
		ToAddresses: []*string{
			aws.String(recipient),
		},
	}

	if len(datas) != 0 {
		return sendEmailWithTemplateData(sesClient, dest, sender, templateName, datas[0])
	}
	return sendEmail(sesClient, dest, sender, templateName)
}

func sendEmailWithTemplateData(sesClient *sessdk.SES, dest *ses.Destination, source, templName string, datas map[string]interface{}) error {
	templateDatas, err := genTemplateDatas(datas)
	if err != nil {
		return err
	}

	templatedInput := &ses.SendTemplatedEmailInput{
		Destination:  dest,
		Source:       aws.String(source),
		Template:     aws.String(templName),
		TemplateData: templateDatas,
	}
	_, err = sesClient.SendTemplatedEmail(templatedInput)

	if err != nil {
		return err
	}
	return nil
}

func sendEmail(sesClient *sessdk.SES, dest *ses.Destination, source, templName string) error {
	data, err := getSEStemplate(sesClient, aws.String(templName))
	if err != nil {
		return err
	}

	input := &ses.SendEmailInput{
		Destination: dest,
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(charSet),
					Data:    aws.String(*data.Template.HtmlPart),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(charSet),
				Data:    aws.String(*data.Template.SubjectPart),
			},
		},
		Source: aws.String(source),
	}

	_, err = sesClient.SendEmail(input)
	if err != nil {
		return err
	}

	return nil
}

func genTemplateDatas(datas map[string]interface{}) (*string, error) {
	rawJSON, marshalErr := json.Marshal(datas)
	if marshalErr != nil {
		return nil, marshalErr
	}

	jsonStr := string(rawJSON)

	return &jsonStr, nil
}
