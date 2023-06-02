/*
 * @Author: xx
 * @Date: 2023-04-24 10:53:04
 * @LastEditTime: 2023-05-15 20:24:32
 * @Description:
 */
package response

type PageResult struct {
	List          interface{} `json:"list"`
	SearchOptions interface{} `json:"searchOptions"`
	Total         int64       `json:"total"`
	Page          int         `json:"page"`
	PageSize      int         `json:"pageSize"`
}
