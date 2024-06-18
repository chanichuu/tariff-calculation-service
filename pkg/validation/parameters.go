package validation

type PartitionId struct {
	PartitionId string `uri:"partitionId" binding:"required,uuid4"`
}

type PartitionIdWithId struct {
	PartitionId string `uri:"partitionId" binding:"required,uuid4"`
	Id          string `uri:"id" binding:"required,uuid4"`
}
