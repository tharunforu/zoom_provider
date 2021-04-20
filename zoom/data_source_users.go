package zoom

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUsers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceUsersRead,
		Schema: map[string]*schema.Schema{

			"users": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"first_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"email": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}
func dataSourceUsersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}
	bearer := "Bearer " + "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MTg5MzgzMjMsImlhdCI6MTYxODg1MTkyNH0.ngd_dOTYMp5ftwP2W-R8XpxHU1dX0i2o6B5xslwLDJ8"
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users", "https://api.zoom.us/v2"), nil)
	if err != nil {
		return diag.FromErr(err)
	}
	req.Header.Add("Authorization", bearer)
	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	whole_body := make([]map[string]interface{}, 0)
	err = json.NewDecoder(r.Body).Decode(&whole_body)
	if err != nil {
		return diag.FromErr(err)
	}
	//usersli:=flatternUsers(&users)
	if whole_body.users != nil {
		ois := make([]interface{}, len(*orderItems), len(*orderItems))

		for i, orderItem := range *orderItems {
			oi := make(map[string]interface{})

			oi["id"] = orderItem.Coffee.ID
			oi["first_name"] = orderItem.Coffee.Name
			oi["last_name"] = orderItem.Coffee.Teaser
			oi["email"] = orderItem.Coffee.Description

			ois[i] = oi
		}

		return ois
	}

	return make([]interface{}, 0)

	if err := d.Set("users", usersli); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}