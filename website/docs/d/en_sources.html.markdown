---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_sources'
description: |-
  List all the destinations
---

# ibm_en_sources

Provides a read-only data source for sources. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example usage

```terraform
data "ibm_en_sources" "en_sources" {
  instance_guid = ibm_resource_instance.en_terraform_test_resource.guid
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `search_key` - (Optional, String) Filter the destinations by name or type.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the `en_sources`.

- `sources` - (List) List of destinations.

  - `id` - (String) Source ID.

  - `name` - (String) Source name.

  - `description` - (String) Source description.

  - `topic_count` - (Integer) Subscription count.

  - `enabled` - (bool) Source is enabled or not.

  - `type` - (String) source type api/IBM Cloud Monitoring/resource lifecycle events/ App Configuration/ Secrets Manager/ SCC.

  - `updated_at` - (String) Lats updated time.

- `total_count` - (Integer) Total number of destinations.
