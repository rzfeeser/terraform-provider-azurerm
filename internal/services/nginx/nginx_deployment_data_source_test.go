// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nginx_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
)

type NginxDeploymentDataSource struct{}

func TestAccNginxDeploymentDataSource_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_nginx_deployment", "test")
	r := NginxDeploymentDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("location").Exists(),
				check.That(data.ResourceName).Key("nginx_version").Exists(),
				check.That(data.ResourceName).Key("sku").Exists(),
				check.That(data.ResourceName).Key("capacity").Exists(),
				check.That(data.ResourceName).Key("managed_resource_group").Exists(),
				check.That(data.ResourceName).Key("ip_address").Exists(),
			),
		},
	})
}

func (d NginxDeploymentDataSource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

data "azurerm_nginx_deployment" "test" {
  name                = azurerm_nginx_deployment.test.name
  resource_group_name = azurerm_nginx_deployment.test.resource_group_name
}
`, DeploymentResource{}.basic(data))
}
