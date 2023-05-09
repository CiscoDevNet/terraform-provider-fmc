package fmc

import (
    "context"

    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFmcStaticIPv4Route() *schema.Resource {
    return &schema.Resource{
        Description: "Data source for Static IPv4 route in FMC\n\n" +
            "An example is shown below: \n" +
            "```hcl\n" +
            "data \"fmc_staticIPv4_route\" \"route\" {\n" +
            "   device_id = \"<device ID>\"\n" +
            "}\n" +
            "```",
        ReadContext: dataSourceFmcStaticIPv4RouteRead,
        Schema: map[string]*schema.Schema{
            "device_id": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "The ID of this resource",
            },
            "network_name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "The ID of this resource",
            },
            "type": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: "Type of this resource",
            },
        },
    }
}

func dataSourceFmcStaticIPv4RouteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    c := m.(*Client)

    // Warning or errors can be collected in a slice type
    var diags diag.Diagnostics
    staticIpv4Route, err := c.GetStaticIPv4RouteByName(ctx, d.Get("device_id").(string), d.Get("network_name").(string))

    if err != nil {
        diags = append(diags, diag.Diagnostic{
            Severity: diag.Error,
            Summary:  "unable to get static ipv4 route",
            Detail:   err.Error(),
        })
        return diags
    }

    d.SetId(staticIpv4Route.ID)

    if err := d.Set("network_name", staticIpv4Route.InterfaceName); err != nil {
        diags = append(diags, diag.Diagnostic{
            Severity: diag.Error,
            Summary:  "unable to get static ipv4 route",
            Detail:   err.Error(),
        })
        return diags
    }

    if err := d.Set("type", staticIpv4Route.Type); err != nil {
        diags = append(diags, diag.Diagnostic{
            Severity: diag.Error,
            Summary:  "unable to get static ipv4 route",
            Detail:   err.Error(),
        })
        return diags
    }

    return diags
}
