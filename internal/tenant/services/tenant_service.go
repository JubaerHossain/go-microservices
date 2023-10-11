package services

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/JubaerHossain/go-service/internal/tenant/models"
	"github.com/JubaerHossain/go-service/internal/tenant/validation"
	pagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AllTenant(requestFilter map[string]interface{}) ([]models.Tenant, pagination.PaginationData) {
	var tenants []models.Tenant

	filter := bson.M{}

	if requestFilter["status"] != "" {
		filter["status"] = requestFilter["status"]
	}

	page, _ := strconv.ParseInt(requestFilter["page"].(string), 10, 64)
	limit, _ := strconv.ParseInt(requestFilter["limit"].(string), 10, 64)

	paginatedData, err := pagination.New(models.TenantCollection.Collection).
		Page(page).
		Limit(limit).
		Sort("created_at", -1).
		Decode(&tenants).
		Filter(filter).
		Find()

	if err != nil {
		panic(err)
	}
	return tenants, paginatedData.Pagination
}

func CreateATenant(createTenant validation.CreateTenantRequest) models.Tenant {
	tenant := models.Tenant{
		Id:        primitive.NewObjectID(),
		Task:      createTenant.Task,
		Status:    createTenant.Status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := models.TenantCollection.InsertOne(tenant)
	if err != nil || result == nil {
		panic(err)
	}

	return tenant
}

func UpdateATenant(tenantId string, updateTenant validation.UpdateTenantRequest) (models.Tenant, error) {

	objId, _ := primitive.ObjectIDFromHex(tenantId)

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	result := models.TenantCollection.FindOneAndUpdate(
		bson.M{"_id": objId},
		bson.D{
			{"$set", bson.M{
				"task":       updateTenant.Task,
				"status":     updateTenant.Status,
				"updated_at": time.Now(),
			}},
		},
		&opt,
	)

	if result.Err() != nil {
		log.Println("Err ", result.Err())
		return models.Tenant{}, result.Err()
	}

	var tenant models.Tenant
	if err := result.Decode(&tenant); err != nil {
		return models.Tenant{}, err
	}

	return tenant, nil
}

func ATenant(tenantId string) models.Tenant {
	var tenant models.Tenant

	objId, _ := primitive.ObjectIDFromHex(tenantId)

	err := models.TenantCollection.FindOne(bson.M{"_id": objId}).Decode(&tenant)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return tenant
}

func DeleteATenant(tenantId string) bool {
	objId, _ := primitive.ObjectIDFromHex(tenantId)

	result := models.TenantCollection.FindOneAndDelete(bson.D{{"_id", objId}})

	if result.Err() != nil {
		return false
	}

	return true
}
