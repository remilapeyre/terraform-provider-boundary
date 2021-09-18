// Code generated by scripts/generate_datasource.go. DO NOT EDIT.
//go:generate go run ../../scripts/generate_datasource.go -name AuthMethods -path auth-methods

// This file was generated based on Boundary v0.6.1

package provider

import (
	"context"
	"net/url"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var dataSourceAuthMethodsSchema = map[string]*schema.Schema{
	"filter": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"items": {
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"authorized_actions": {
					Type:        schema.TypeList,
					Description: "Output only. The available actions on this resource for this user.",
					Computed:    true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
				"created_time": {
					Type:        schema.TypeString,
					Description: "Output only. The time this resource was created.",
					Computed:    true,
				},
				"description": {
					Type:        schema.TypeString,
					Description: "Optional user-set description for identification purposes.",
					Computed:    true,
				},
				"id": {
					Type:        schema.TypeString,
					Description: "Output only. The ID of the Auth Method.",
					Computed:    true,
				},
				"is_primary": {
					Type:        schema.TypeBool,
					Description: "Output only. Whether this auth method is the primary auth method for it's scope.\nTo change this value update the primary_auth_method_id field on the scope.",
					Computed:    true,
				},
				"name": {
					Type:        schema.TypeString,
					Description: "Optional name for identification purposes.",
					Computed:    true,
				},
				"scope": {
					Type:     schema.TypeList,
					Computed: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"description": {
								Type:        schema.TypeString,
								Description: "Output only. The description of the Scope, if any.",
								Computed:    true,
							},
							"id": {
								Type:        schema.TypeString,
								Description: "Output only. The ID of the Scope.",
								Computed:    true,
							},
							"name": {
								Type:        schema.TypeString,
								Description: "Output only. The name of the Scope, if any.",
								Computed:    true,
							},
							"parent_scope_id": {
								Type:        schema.TypeString,
								Description: "Output only. The ID of the parent Scope, if any. This field will be empty if this is the \"global\" scope.",
								Computed:    true,
							},
							"type": {
								Type:        schema.TypeString,
								Description: "Output only. The type of the Scope.",
								Computed:    true,
							},
						},
					},
				},
				"scope_id": {
					Type:        schema.TypeString,
					Description: "The ID of the Scope of which this Auth Method is a part.",
					Computed:    true,
				},
				"type": {
					Type:        schema.TypeString,
					Description: "The Auth Method type.",
					Computed:    true,
				},
				"updated_time": {
					Type:        schema.TypeString,
					Description: "Output only. The time this resource was last updated.",
					Computed:    true,
				},
				"version": {
					Type:        schema.TypeInt,
					Description: "Version is used in mutation requests, after the initial creation, to ensure this resource has not changed.\nThe mutation will fail if the version does not match the latest known good version.",
					Computed:    true,
				},
			},
		},
	},
	"recursive": {
		Type:     schema.TypeBool,
		Optional: true,
	},
	"scope_id": {
		Type:     schema.TypeString,
		Optional: true,
	},
}

func dataSourceAuthMethods() *schema.Resource {
	return &schema.Resource{
		Description: "Lists all Auth Methods.",
		Schema:      dataSourceAuthMethodsSchema,
		ReadContext: dataSourceAuthMethodsRead,
	}
}

func dataSourceAuthMethodsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*metaData).client

	req, err := client.NewRequest(ctx, "GET", "auth-methods", nil)
	if err != nil {
		return diag.FromErr(err)
	}

	q := url.Values{}
	q.Add("filter", d.Get("filter").(string))
	recursive := d.Get("recursive").(bool)
	if recursive {
		q.Add("recursive", strconv.FormatBool(recursive))
	}
	q.Add("scope_id", d.Get("scope_id").(string))
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		diag.FromErr(err)
	}
	apiError, err := resp.Decode(nil)
	if err != nil {
		return diag.FromErr(err)
	}
	if apiError != nil {
		return apiErr(apiError)
	}
	err = set(dataSourceAuthMethodsSchema, d, resp.Map)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("boundary-auth-methods")

	return nil
}
