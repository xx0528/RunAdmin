/*
 * @Author: xx
 * @Date: 2023-04-24 10:53:04
 * @LastEditTime: 2023-05-18 10:47:14
 * @Description:
 */
package config

type Local struct {
	Path      string `mapstructure:"path" json:"path" yaml:"path"`                   // 本地文件访问路径
	StorePath string `mapstructure:"store-path" json:"store-path" yaml:"store-path"` // 本地文件存储路径
}
