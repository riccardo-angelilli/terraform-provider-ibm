---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_cloud_connections"
description: |-
  Manages IBM cloud_connections in the Power Virtual Server cloud.
---

# ibm_pi_cloud_connections

Retrieve information about all cloud connections as a read-only data source. For more information, about IBM power virtual server cloud, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example Usage

```terraform
data "ibm_pi_cloud_connections" "example" {
  pi_cloud_instance_id      = "<value of the cloud_instance_id>"
}
```

### Notes

- Please find [supported Regions](https://cloud.ibm.com/apidocs/power-cloud#endpoint) for endpoints.
- If a Power cloud instance is provisioned at `lon04`, The provider level attributes should be as follows:

  - `region` - `lon`
  - `zone` - `lon04`

Example usage:

  ```terraform
    provider "ibm" {
      region    =   "lon"
      zone      =   "lon04"
    }
  ```

## Argument Reference

Review the argument references that you can specify for your data source.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `connections` - (List) List of all the Cloud Connections.

  Nested scheme for `connections`:
  - `classic_enabled` - (Boolean) Enable classic endpoint destination.
  - `cloud_connection_id` - (String) The unique identifier of the cloud connection.
  - `connection_mode` - (String) Type of service the gateway is attached to.
  - `global_routing` - (String) Enable global routing for this cloud connection.
  - `gre_destination_address` - (String) GRE destination IP address.
  - `gre_source_address` - (String) GRE auto-assigned source IP address.
  - `ibm_ip_address` - (String) IBM IP address.
  - `metered` - (String) Enable metering for this cloud connection.
  - `name` - (String) Name of the cloud connection.
  - `networks` - (Set) Set of Networks attached to this cloud connection.
  - `port` - (String) Port.
  - `speed` - (Integer) Speed of the cloud connection (speed in megabits per second).
  - `status` - (String) Link status.
  - `user_ip_address` - (String) User IP address.
  - `vpc_crns` - (Set) Set of VPCs attached to this cloud connection.
  - `vpc_enabled` - (Boolean) Enable VPC for this cloud connection.
