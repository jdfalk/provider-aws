/*
Copyright 2020 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package integration

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/apigatewayv2"

	svcapitypes "github.com/crossplane/provider-aws/apis/apigatewayv2/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.
// TODO(muvaf): We can generate one-time boilerplate for these hooks but currently
// ACK doesn't support not generating if file exists.

// GenerateGetIntegrationInput returns input for read
// operation.
func GenerateGetIntegrationInput(cr *svcapitypes.Integration) *svcsdk.GetIntegrationInput {
	res := preGenerateGetIntegrationInput(cr, &svcsdk.GetIntegrationInput{})

	if cr.Status.AtProvider.IntegrationID != nil {
		res.SetIntegrationId(*cr.Status.AtProvider.IntegrationID)
	}

	return postGenerateGetIntegrationInput(cr, res)
}

// GenerateIntegration returns the current state in the form of *svcapitypes.Integration.
func GenerateIntegration(resp *svcsdk.GetIntegrationOutput) *svcapitypes.Integration {
	cr := &svcapitypes.Integration{}

	if resp.ApiGatewayManaged != nil {
		cr.Status.AtProvider.APIGatewayManaged = resp.ApiGatewayManaged
	}
	if resp.IntegrationId != nil {
		cr.Status.AtProvider.IntegrationID = resp.IntegrationId
	}
	if resp.IntegrationResponseSelectionExpression != nil {
		cr.Status.AtProvider.IntegrationResponseSelectionExpression = resp.IntegrationResponseSelectionExpression
	}

	return cr
}

// GenerateCreateIntegrationInput returns a create input.
func GenerateCreateIntegrationInput(cr *svcapitypes.Integration) *svcsdk.CreateIntegrationInput {
	res := preGenerateCreateIntegrationInput(cr, &svcsdk.CreateIntegrationInput{})

	if cr.Spec.ForProvider.ConnectionID != nil {
		res.SetConnectionId(*cr.Spec.ForProvider.ConnectionID)
	}
	if cr.Spec.ForProvider.ConnectionType != nil {
		res.SetConnectionType(*cr.Spec.ForProvider.ConnectionType)
	}
	if cr.Spec.ForProvider.ContentHandlingStrategy != nil {
		res.SetContentHandlingStrategy(*cr.Spec.ForProvider.ContentHandlingStrategy)
	}
	if cr.Spec.ForProvider.CredentialsARN != nil {
		res.SetCredentialsArn(*cr.Spec.ForProvider.CredentialsARN)
	}
	if cr.Spec.ForProvider.Description != nil {
		res.SetDescription(*cr.Spec.ForProvider.Description)
	}
	if cr.Spec.ForProvider.IntegrationMethod != nil {
		res.SetIntegrationMethod(*cr.Spec.ForProvider.IntegrationMethod)
	}
	if cr.Spec.ForProvider.IntegrationSubtype != nil {
		res.SetIntegrationSubtype(*cr.Spec.ForProvider.IntegrationSubtype)
	}
	if cr.Spec.ForProvider.IntegrationType != nil {
		res.SetIntegrationType(*cr.Spec.ForProvider.IntegrationType)
	}
	if cr.Spec.ForProvider.IntegrationURI != nil {
		res.SetIntegrationUri(*cr.Spec.ForProvider.IntegrationURI)
	}
	if cr.Spec.ForProvider.PassthroughBehavior != nil {
		res.SetPassthroughBehavior(*cr.Spec.ForProvider.PassthroughBehavior)
	}
	if cr.Spec.ForProvider.PayloadFormatVersion != nil {
		res.SetPayloadFormatVersion(*cr.Spec.ForProvider.PayloadFormatVersion)
	}
	if cr.Spec.ForProvider.RequestParameters != nil {
		f11 := map[string]*string{}
		for f11key, f11valiter := range cr.Spec.ForProvider.RequestParameters {
			var f11val string
			f11val = *f11valiter
			f11[f11key] = &f11val
		}
		res.SetRequestParameters(f11)
	}
	if cr.Spec.ForProvider.RequestTemplates != nil {
		f12 := map[string]*string{}
		for f12key, f12valiter := range cr.Spec.ForProvider.RequestTemplates {
			var f12val string
			f12val = *f12valiter
			f12[f12key] = &f12val
		}
		res.SetRequestTemplates(f12)
	}
	if cr.Spec.ForProvider.TemplateSelectionExpression != nil {
		res.SetTemplateSelectionExpression(*cr.Spec.ForProvider.TemplateSelectionExpression)
	}
	if cr.Spec.ForProvider.TimeoutInMillis != nil {
		res.SetTimeoutInMillis(*cr.Spec.ForProvider.TimeoutInMillis)
	}
	if cr.Spec.ForProvider.TLSConfig != nil {
		f15 := &svcsdk.TlsConfigInput{}
		if cr.Spec.ForProvider.TLSConfig.ServerNameToVerify != nil {
			f15.SetServerNameToVerify(*cr.Spec.ForProvider.TLSConfig.ServerNameToVerify)
		}
		res.SetTlsConfig(f15)
	}

	return postGenerateCreateIntegrationInput(cr, res)
}

// GenerateDeleteIntegrationInput returns a deletion input.
func GenerateDeleteIntegrationInput(cr *svcapitypes.Integration) *svcsdk.DeleteIntegrationInput {
	res := preGenerateDeleteIntegrationInput(cr, &svcsdk.DeleteIntegrationInput{})

	if cr.Status.AtProvider.IntegrationID != nil {
		res.SetIntegrationId(*cr.Status.AtProvider.IntegrationID)
	}

	return postGenerateDeleteIntegrationInput(cr, res)
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "NotFoundException"
}
