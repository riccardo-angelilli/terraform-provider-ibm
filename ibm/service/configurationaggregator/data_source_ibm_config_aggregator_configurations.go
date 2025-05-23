// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.92.0-af5c89a5-20240617-153232
 */

package configurationaggregator

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/configuration-aggregator-go-sdk/configurationaggregatorv1"
)

func DataSourceIbmConfigAggregatorConfigurations() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmConfigAggregatorConfigurationsRead,

		Schema: map[string]*schema.Schema{
			"config_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of resource configuration that are to be retrieved.",
			},
			"service_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the IBM Cloud service for which resources are to be retrieved.",
			},
			"resource_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource group id of the resources.",
			},
			"location": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The location or region in which the resources are created.",
			},
			"resource_crn": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The crn of the resource.",
			},
			"sub_account": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Filter the resource configurations from the specified sub-account in an enterprise hierarchy.",
				ValidateFunc: validate.InvokeDataSourceValidator("ibm_config_aggregator_configurations", "sub_account"),
			},
			"access_tags": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Filter the resource configurations attached with the specified access tags.",
				ValidateFunc: validate.InvokeDataSourceValidator("ibm_config_aggregator_configurations", "access_tags"),
			},
			"user_tags": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Filter the resource configurations attached with the specified user tags.",
				ValidateFunc: validate.InvokeDataSourceValidator("ibm_config_aggregator_configurations", "user_tags"),
			},
			"service_tags": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Filter the resource configurations attached with the specified service tags.",
				ValidateFunc: validate.InvokeDataSourceValidator("ibm_config_aggregator_configurations", "service_tags"),
			},
			"prev": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The reference to the previous page of entries.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The reference to the previous page of entries.",
						},
						"start": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "the start string for the query to view the page.",
						},
					},
				},
			},
			"configs": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Array of resource configurations.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"about": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The basic metadata fetched from the query API.",
						},
						"config": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The configuration of the resource.",
						},
					},
				},
			},
		},
	}
}
func DataSourceIbmConfigAggregatorValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 4)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "sub_account",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Regexp:                     `^[a-zA-Z0-9]+$`,
			MinValueLength:             32,
			MaxValueLength:             32,
		},
		validate.ValidateSchema{
			Identifier:                 "access_tags",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Regexp:                     `^[a-zA-Z0-9:-]+$`,
			MinValueLength:             1,
			MaxValueLength:             128,
		},
		validate.ValidateSchema{
			Identifier:                 "user_tags",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Regexp:                     `^[a-zA-Z0-9:-]+$`,
			MinValueLength:             1,
			MaxValueLength:             128,
		},
		validate.ValidateSchema{
			Identifier:                 "service_tags",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Regexp:                     `^[a-zA-Z0-9:-]+$`,
			MinValueLength:             1,
			MaxValueLength:             128,
		},
	)
	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_config_aggregator_configurations", Schema: validateSchema}
	return &resourceValidator
}

func dataSourceIbmConfigAggregatorConfigurationsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	configurationAggregatorClient, err := meta.(conns.ClientSession).ConfigurationAggregatorV1()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_config_aggregator_configurations", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	region := getConfigurationInstanceRegion(configurationAggregatorClient, d)
	instanceId := d.Get("instance_id").(string)
	log.Printf("Fetching config for instance_id: %s", instanceId)
	configurationAggregatorClient = getClientWithConfigurationInstanceEndpoint(configurationAggregatorClient, instanceId, region)

	listConfigsOptions := &configurationaggregatorv1.ListConfigsOptions{}

	if _, ok := d.GetOk("config_type"); ok {
		listConfigsOptions.SetConfigType(d.Get("config_type").(string))
	}
	if _, ok := d.GetOk("service_name"); ok {
		listConfigsOptions.SetServiceName(d.Get("service_name").(string))
	}
	if _, ok := d.GetOk("resource_group_id"); ok {
		listConfigsOptions.SetResourceGroupID(d.Get("resource_group_id").(string))
	}
	if _, ok := d.GetOk("location"); ok {
		listConfigsOptions.SetLocation(d.Get("location").(string))
	}
	if _, ok := d.GetOk("resource_crn"); ok {
		listConfigsOptions.SetResourceCrn(d.Get("resource_crn").(string))
	}
	if _, ok := d.GetOk("sub_account"); ok {
		listConfigsOptions.SetSubAccount(d.Get("sub_account").(string))
	}
	if _, ok := d.GetOk("access_tags"); ok {
		listConfigsOptions.SetAccessTags(d.Get("access_tags").(string))
	}
	if _, ok := d.GetOk("user_tags"); ok {
		listConfigsOptions.SetUserTags(d.Get("user_tags").(string))
	}
	if _, ok := d.GetOk("service_tags"); ok {
		listConfigsOptions.SetServiceTags(d.Get("service_tags").(string))
	}

	var pager *configurationaggregatorv1.ConfigsPager
	pager, err = configurationAggregatorClient.NewConfigsPager(listConfigsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_config_aggregator_configurations", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	allItems, err := pager.GetAll()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ConfigsPager.GetAll() failed %s", err), "(Data) ibm_config_aggregator_configurations", "read")
		log.Printf("[DEBUG] %s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmConfigAggregatorConfigurationsID(d))

	mapSlice := []map[string]interface{}{}
	for _, model := range allItems {
		modelMap, err := DataSourceIbmConfigAggregatorConfigurationsConfigToMap(&model)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_config_aggregator_configurations", "read")
			return tfErr.GetDiag()
		}
		mapSlice = append(mapSlice, modelMap)
	}

	if err = d.Set("configs", mapSlice); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting configs %s", err), "(Data) ibm_config_aggregator_configurations", "read")
		return tfErr.GetDiag()
	}

	return nil
}

// dataSourceIbmConfigAggregatorConfigurationsID returns a reasonable ID for the list.
func dataSourceIbmConfigAggregatorConfigurationsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmConfigAggregatorConfigurationsPaginatedPreviousToMap(model *configurationaggregatorv1.PaginatedPrevious) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	if model.Start != nil {
		modelMap["start"] = *model.Start
	}
	return modelMap, nil
}

func DataSourceIbmConfigAggregatorConfigurationsConfigToMap(model *configurationaggregatorv1.Config) (map[string]interface{}, error) {
	if model == nil {
		return nil, fmt.Errorf("model is nil")
	}

	// Convert About struct to JSON string or "{}"
	aboutStr := "{}"
	if model.About != nil {
		bytes, err := json.Marshal(model.About)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal About: %v", err)
		}
		aboutStr = string(bytes)
	}

	// Convert Config struct to JSON string or "{}"
	configStr := "{}"
	if model.Config != nil {
		bytes, err := json.Marshal(model.Config)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal Config: %v", err)
		}
		configStr = string(bytes)
	}

	return map[string]interface{}{
		"about":  aboutStr,
		"config": configStr,
	}, nil
}

func DataSourceIbmConfigAggregatorConfigurationsAboutToMap(model *configurationaggregatorv1.About) (string, error) {
	modelMap := make(map[string]interface{})
	modelMap["account_id"] = *model.AccountID
	modelMap["config_type"] = *model.ConfigType
	modelMap["resource_crn"] = *model.ResourceCrn
	modelMap["resource_group_id"] = *model.ResourceGroupID
	modelMap["service_name"] = *model.ServiceName
	modelMap["resource_name"] = *model.ResourceName
	modelMap["last_config_refresh_time"] = model.LastConfigRefreshTime.String()
	modelMap["location"] = *model.Location
	modelMap["access_tags"] = model.AccessTags
	modelMap["user_tags"] = model.UserTags
	modelMap["service_tags"] = model.ServiceTags
	jsonData, err := json.Marshal(modelMap)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
	// modelMap["tags"] = make(map[string]interface{})
}

// func DataSourceIbmConfigAggregatorConfigurationsTagsToMap(model *configurationaggregatorv1.Tags) (map[string]interface{}, error) {
// 	modelMap := make(map[string]interface{})
// 	if model.Tag != nil {
// 		modelMap["tag"] = *model.Tag
// 	}
// 	return modelMap, nil
// }

func DataSourceIbmConfigAggregatorConfigurationsConfigurationToMap(model *configurationaggregatorv1.Configuration) (string, error) {
	if model == nil {
		return "", nil
	}

	if len(model.GetProperties()) != 0 {
		checkMap := model.GetProperties()
		tryMap := make(map[string]interface{})
		for i, v := range checkMap {
			tryMap[i] = v
		}
		jsonData, err := json.Marshal(tryMap)
		if err != nil {
			return "", err
		}
		return string(jsonData), nil
	}
	return "", nil
}
