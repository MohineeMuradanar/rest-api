package controller

import (
	"AssDeploy/model"
	"AssDeploy/utility"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection = utility.DB()

// swagger:operation POST /students Student createNewStudent
//
// Add new Student
//
// Returns new Student
//
// ---
// produces:
// - application/json
// parameters:
// - name: student
//   in: body
//   description: add Student data
//   required: true
//   schema:
//     "$ref": "#/definitions/Student"
// responses:
//   '200':
//     description:  New Student created
//     schema:
//       "$ref": "#/definitions/Student"

//createStudent
func createNewStudent(resp http.ResponseWriter, req *http.Request) {

	fmt.Println("Test", req.Body)
	var student model.Student
	fmt.Println("Test", req.Body)
	ctx := req.Context()
	err := json.NewDecoder(req.Body).Decode(&student)
	if err != nil {
		log.Println(err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	// collection := client.Database("University").Collection("studentdata")
	// collection := utility.DB()

	result, err := collection.InsertOne(ctx, student)
	if err != nil {
		log.Println(err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("insertion done", result)
	log.Println(result)

	id , ok := result.InsertedID.(primitive.ObjectID)
	log.Println(id ,ok )
	student.ID = id

	resp.Header().Add("content-type", "appllication/json")

	json.NewEncoder(resp).Encode(student)

}

// swagger:operation GET /students/{id} Student getById
//
// Get Student
//
// Returns existing Student filtered by id
//
// ---
// produces:
// - application/json
// parameters:
//  - name: id
//    type: string
//    in: path
//    required: true
// responses:
//   '200':
//     description: Student response
//     schema:
//      "$ref": "#/definitions/Student"

//GetStudentbyid fetch Student from db
func getById(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("content-type", "application/json")
	var student model.Student

	param := mux.Vars(req)
	fmt.Println("mux var value", param)
	id, err := primitive.ObjectIDFromHex(param["id"])
	if err != nil {
		log.Println(err)
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	filter := model.Student{ID: id}

	// filter := bson.M{"_id": id}
	// collection := client.Database("University").Collection("studentdata")
	ctx := req.Context()

	// collection := utility.DB()

	err = collection.FindOne(ctx, filter).Decode(&student)
	if err != nil {
		log.Println(err)
		resp.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Printf("%+v\n", student)

	json.NewEncoder(resp).Encode(student)
}

// swagger:operation GET /students Student getAll
//
// Get Student
//
// Returns existing student
//
// ---
// produces:
// - application/json
// responses:
//   '200':
//     description: Student response
//     schema:
//      "$ref": "#/definitions/Student"

//GetStudent fetch all  from db
func getAll(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("In get method \n ")
	var studentinfo []model.Student
	name := req.URL.Query().Get("name")
	city := req.URL.Query().Get("city")
	country := req.URL.Query().Get("country")

	// value := []primitive.M{}
	filter := primitive.M{}
	ctx := req.Context()
	// collection := client.Database("University").Collection("studentdata")
	// collection := utility.DB()

	if name != "" {
		filter["name"] = name
		// value = append(value, primitive.M{"name": name})
	}

	if city != "" {
		filter["city"] = city

		// value = append(value, primitive.M{"city": city})
	}

	if country != "" {
		filter["country"] = country

		// value = append(value, primitive.M{"country": country})
	}

	// if len(value) > 0 {
	// 	filter = primitive.M{"$and": value}
	// }
	// primitive.M{ }

	fmt.Println(name, city, country)

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		log.Println(err)
		resp.WriteHeader(http.StatusOK)
		return
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var student model.Student
		err := cur.Decode(&student)
		if err != nil {
			log.Panicln(err)
			resp.WriteHeader(http.StatusNotFound)
			return
		}
		studentinfo = append(studentinfo, student)
	}

	// fmt.Printf("%+v\n", studentinfo)
	resp.Header().Add("content-type", "application/json")

	json.NewEncoder(resp).Encode(studentinfo)

}

// swagger:operation PATCH /students/{id} Student updateStudent_Patch
//
// Update Student
//
// Update existing Student filtered by id
//
// ---
// consumes:
// - application/json
// produces:
// - application/json
// parameters:
// - name: id
//   type: string
//   in: path
//   required: true
// - name: student
//   in: body
//   description: add student data
//   required: true
//   schema:
//     "$ref": "#/definitions/Student"
// responses:
//   '200':
//     description: Student response
//     schema:
//       "$ref": "#/definitions/Student"

//updateStudent_Patch update student
func updateStudent_Patch(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("in update method")
	resp.Header().Add("content-type", "application/json")
	var student model.Student
	json.NewDecoder(req.Body).Decode(&student)

	param := mux.Vars(req)

	id, err := primitive.ObjectIDFromHex(param["id"])
	if err != nil {
		log.Println(err)
		resp.WriteHeader(http.StatusBadRequest)
		return

	}

	filter := bson.M{"_id": id}

	update := primitive.M{"$set": student}

	// collection := client.Database("University").Collection("studentdata")
	// collection := utility.DB()

	ctx := req.Context()

	err = collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&student)
	if err != nil {
		log.Println(err)
		resp.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(resp).Encode(student)

}

// swagger:operation PUT /students/{id} Student replaceStudent_Put
//
// replace Student
//
// replace existing Student filtered by id
//
// ---
// consumes:
// - application/json
// produces:
// - application/json
// parameters:
// - name: id
//   type: string
//   in: path
//   required: true
// - name: student
//   in: body
//   description: add student data
//   required: true
//   schema:
//     "$ref": "#/definitions/Student"
// responses:
//   '200':
//     description: Student response
//     schema:
//       "$ref": "#/definitions/Student"

//updateStudent_Patch update student
func replaceStudent_Put(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("in replace method")
	resp.Header().Add("content-type", "application/json")
	var student model.Student
	err:= json.NewDecoder(req.Body).Decode(&student)
	if err!= nil{
		log.Println(err)
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	param := mux.Vars(req)

	id, err := primitive.ObjectIDFromHex(param["id"])
	if err != nil {
		log.Println(err)
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	filter := bson.M{"_id": id}

	update := primitive.M{
		"name":            student.Name,
		"city":            student.City,
		"country":         student.Country,
		"course":          student.Course,
		"yearofadmission": student.YearOfAdmission}

	// collection := client.Database("University").Collection("studentdata")
	// collection := utility.DB()

	ctx := req.Context()

	// err = collection.FindOneAndReplace(ctx, filter, bson.M{"$set": update}, options.FindOneAndReplace().SetReturnDocument(options.After)).Decode(&student)

	err = collection.FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&student)
	if err != nil {
		log.Println(err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(resp).Encode(student)

}

// swagger:operation DELETE /students/{id} Student deleteStudent
//
// Delete  student
//
// Delete existing Student filtered by id
//
// ---
//
// parameters:
//  - name: id
//    type: string
//    in: path
//    required: true
// responses:
//   '200':
//     description: delete Student sucessfully

//Deletestudent fetch Student from db
func deleteStudent(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("in delete method")

	resp.Header().Add("content-type", "application/json")
	// collection := client.Database("University").Collection("studentdata")
	// collection := utility.DB()
	var student model.Student
	err:= json.NewDecoder(req.Body).Decode(&student)
	if err!= nil{
		log.Println(err)
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := req.Context()
	params := mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Println(err)
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	// filter := bson.M{"_id": id}
	deleteresult, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Println(err)
		resp.WriteHeader(http.StatusGone)
		return
	}
	log.Println(deleteresult)

	student.ID = id

	
	json.NewEncoder(resp).Encode(student)
}

func home(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "Hello World!")
}

func check(resp http.ResponseWriter, req *http.Request) {
	os.Exit(0)

}
