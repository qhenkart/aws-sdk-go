package cloudsearchdomainiface

import (
	"github.com/awslabs/aws-sdk-go/service/cloudsearchdomain"
)

type CloudSearchDomainAPI interface {
	Search(*cloudsearchdomain.SearchInput) (*cloudsearchdomain.SearchOutput, error)

	Suggest(*cloudsearchdomain.SuggestInput) (*cloudsearchdomain.SuggestOutput, error)

	UploadDocuments(*cloudsearchdomain.UploadDocumentsInput) (*cloudsearchdomain.UploadDocumentsOutput, error)
}