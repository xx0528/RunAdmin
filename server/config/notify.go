/*
 * @Author: xx
 * @Date: 2023-05-08 16:13:38
 * @LastEditTime: 2023-05-09 11:04:53
 * @Description:
 */
package config

type Notify struct {
	Url    string `mapstructure:"url" json:"url" yaml:"url"`          // url
	Token  string `mapstructure:"token" json:"token" yaml:"token"`    // token
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"` // 密钥
}
