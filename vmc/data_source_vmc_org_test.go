package vmc

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceVmcOrg_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceVmcOrgConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.vmc_org.my_org", "display_name", "DX AUTO Core Team"),
					resource.TestCheckResourceAttr("data.vmc_org.my_org", "name", "k0byp9w7"),
				),
			},
		},
	})
}

func testAccDataSourceVmcOrgConfig() string {
	return fmt.Sprintf(`
provider "vmc" {
	refresh_token = %q
}
	
data "vmc_org" "my_org" {
	id = "54937bce-8119-4fae-84f5-e5e066ee90e6"
}
`,
		os.Getenv("REFRESH_TOKEN"),
	)
}