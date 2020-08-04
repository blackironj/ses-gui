package ses

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	sessdk "github.com/aws/aws-sdk-go/service/ses"
)

// ListSEStemplates gets email-templates from AWS-SES
func ListSEStemplates(max int) ([]*sessdk.TemplateMetadata, error) {
	if EamilServiceSess == nil {
		return nil, errors.New("fail to access")
	}
	sesClient := sessdk.New(EamilServiceSess)

	listTemplatesInput := sessdk.ListTemplatesInput{
		MaxItems: aws.Int64(int64(max)),
	}

	listTemplatesOutput, err := sesClient.ListTemplates(&listTemplatesInput)
	if err != nil {
		return nil, err
	}
	return listTemplatesOutput.TemplatesMetadata, nil
}

// UploadSEStemplate uploads a email-template to AWS-SES
func UploadSEStemplate(sesTemplate *sessdk.Template) error {
	sesClient := sessdk.New(EamilServiceSess)

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
	sesClient := sessdk.New(EamilServiceSess)

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
	sesClient := sessdk.New(EamilServiceSess)

	getTemplateInput := &ses.GetTemplateInput{
		TemplateName: name,
	}

	getTemplateOutput, err := sesClient.GetTemplate(getTemplateInput)
	if err != nil {
		return nil, err
	}
	return getTemplateOutput, nil
}

// SendEmailWithUnregisteredTemplate sends a email using unregistered template
func SendEmailWithUnregisteredTemplate(input *sessdk.SendEmailInput) error {
	sesClient := sessdk.New(EamilServiceSess)

	_, err := sesClient.SendEmail(input)
	if err != nil {
		return err
	}
	return nil
}
