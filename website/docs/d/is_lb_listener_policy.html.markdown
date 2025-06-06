---
subcategory: "VPC infrastructure"
page_title: "IBM : ibm_is_lb_listener_policy"
description: |-
  Get information about LoadBalancerListenerPolicy
---

# ibm_is_lb_listener_policy

Provides a read-only data source for LoadBalancerListenerPolicy. For more information, about VPC load balance listener policy, see [monitoring application Load Balancer for VPC metrics](https://cloud.ibm.com/docs/vpc?topic=vpc-monitoring-metrics-alb).

## Example Usage

```terraform
data "ibm_is_lb_listener_policy" "example" {
  lb        = ibm_is_lb.example.id
  listener  = ibm_is_lb_listener.example.listener_id
  policy_id = ibm_is_lb_listener_policy.example.policy_id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

- `policy_id` - (Required, String) The policy identifier.
- `listener` - (Required, String) The listener identifier.
- `lb` - (Required, String) The load balancer identifier.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the LoadBalancerListenerPolicy.
- `action` - (String) The policy action.The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the policy on which the unexpected property value was encountered.

- `created_at` - (String) The date and time that this policy was created.

- `href` - (String) The listener policy's canonical URL.

- `name` - (String) The user-defined name for this policy.

- `priority` - (Integer) Priority of the policy. Lower value indicates higher priority.

- `provisioning_status` - (String) The provisioning status of this policy.

- `rules` - (List) The rules for this policy.
Nested scheme for `rules`:
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
	Nested scheme for `deleted`:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The rule's canonical URL.
	- `id` - (String) The rule's unique identifier.

- `target` - (List) -If `action` is `forward_to_pool`, the response is a `LoadBalancerPoolReference`-If `action` is `forward_to_listener`, specify a `LoadBalancerListenerIdentity` in this load balancer to forward to.`- If `action` is `redirect`, the response is a `LoadBalancerListenerPolicyRedirectURL`- If `action` is `https_redirect`, the response is a `LoadBalancerListenerHTTPSRedirect`.
Nested scheme for `target`:
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
	Nested scheme for `deleted`:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The pool's canonical URL.
	- `http_status_code` - (Integer) The HTTP status code for this redirect.
	- `id` - (String) The unique identifier for this load balancer pool.
	- `listener` - (List)
	Nested scheme for `listener`:
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
		Nested scheme for `deleted`:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The listener's canonical URL.
		- `id` - (String) The unique identifier for this load balancer listener.
	- `name` - (String) The user-defined name for this load balancer pool.
	- `uri` - (String) The redirect relative target URI.
	- `url` - (String) The redirect target URL.
