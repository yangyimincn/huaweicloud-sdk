package v1

import (
	"encoding/json"
	"fmt"
)

func (c *CPHClient) UploadStorage(phoneID, bucketName, objPath string, files []string)  {
	uri := fmt.Sprintf("/v1/%s/cloud-phone/phones/batch-storage", c.GetProjectID())

	body := map[string]interface{}{
		"storage_infos": []map[string]interface{}{
			{
				"phone_id":      phoneID,
				"include_files": files,
				"bucket_name":   bucketName,
				"object_path":   objPath,
			},
		},
	}

	bodyB, err := json.Marshal(body)
	if err != nil {
		fmt.Println(err)
	}
	result, err := c.DoRequest("POST", uri, nil, bodyB)
	fmt.Println(string(result))
}
