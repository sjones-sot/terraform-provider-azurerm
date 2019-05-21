package azurerm

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func dataSourceArmRedisCache() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceArmRedisCacheRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"location": locationForDataSourceSchema(),

			"resource_group_name": resourceGroupNameForDataSourceSchema(),

			"zones": singleZonesSchema(),

			"capacity": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"family": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"sku_name": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"minimum_tls_version": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"shard_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"enable_non_ssl_port": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"private_static_ip_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"redis_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"maxclients": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"maxmemory_delta": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						"maxmemory_reserved": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						"maxmemory_policy": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "volatile-lru",
						},

						"maxfragmentationmemory_reserved": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						"rdb_backup_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"rdb_backup_frequency": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"rdb_backup_max_snapshot_count": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"rdb_storage_connection_string": {
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
						},

						"notify_keyspace_events": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"aof_backup_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"aof_storage_connection_string_0": {
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
						},

						"aof_storage_connection_string_1": {
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
						},
						"enable_authentication": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
					},
				},
			},

			"patch_schedule": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"day_of_week": {
							Type:     schema.TypeString,
							Required: true,
						},
						"start_hour_utc": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},

			"hostname": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"port": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"ssl_port": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"primary_access_key": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},

			"secondary_access_key": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},

			"tags": tagsForDataSourceSchema(),
		},
	}
}

func dataSourceArmRedisCacheRead(d *schema.ResourceData, meta interface{}) error {
	ctx := meta.(*ArmClient).StopContext
	client := meta.(*ArmClient).redisClient

	resourceGroup := d.Get("resource_group_name").(string)
	name := d.Get("name").(string)

	resp, err := client.Get(ctx, resourceGroup, name)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			return fmt.Errorf("Error: Redis instance %q (Resource group %q) was not found", name, resourceGroup)
		}
		return fmt.Errorf("Error reading the state of Redis instance %q: %+v", name, err)
	}

	d.SetId(*resp.ID)

	if location := resp.Location; location != nil {
		d.Set("location", azureRMNormalizeLocation(*location))
	}

	if zones := resp.Zones; zones != nil {
		d.Set("zones", zones)
	}

	if sku := resp.Sku; sku != nil {
		d.Set("capacity", sku.Capacity)
		d.Set("family", sku.Family)
		d.Set("sku_name", sku.Name)
	}

	if props := resp.Properties; props != nil {
		d.Set("ssl_port", props.SslPort)
		d.Set("hostname", props.HostName)
		d.Set("minimum_tls_version", string(props.MinimumTLSVersion))
		d.Set("port", props.Port)
		d.Set("enable_non_ssl_port", props.EnableNonSslPort)
		if props.ShardCount != nil {
			d.Set("shard_count", props.ShardCount)
		}
		d.Set("private_static_ip_address", props.StaticIP)
		d.Set("subnet_id", props.SubnetID)
	}

	redisConfiguration, err := flattenRedisConfiguration(resp.RedisConfiguration)
	if err != nil {
		return fmt.Errorf("Error flattening `redis_configuration`: %+v", err)
	}
	if err := d.Set("redis_configuration", redisConfiguration); err != nil {
		return fmt.Errorf("Error setting `redis_configuration`: %+v", err)
	}

	keys, err := client.ListKeys(ctx, resourceGroup, name)
	if err != nil {
		return err
	}

	d.Set("primary_access_key", *keys.PrimaryKey)
	d.Set("secondary_access_key", *keys.SecondaryKey)

	flattenAndSetTags(d, resp.Tags)

	return nil

}