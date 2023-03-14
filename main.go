package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	model "gym/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)


func CustomerEnrolments(w http.ResponseWriter,r *http.Request){


	// params:=mux.Vars(r)
	var newUser model.Users
	json.NewDecoder(r.Body).Decode(&newUser)

	result:=db.Create(&newUser)
	if result.Error!=nil{
		fmt.Println("error in DB")

	}

	fmt.Fprint(w,"Customer Enrolled!!!")

}

func GymEmployEnrolment(w http.ResponseWriter,r *http.Request){

	var GymMember model.GymEmployee
	json.NewDecoder(r.Body).Decode(&GymMember)

	result:=db.Create(&GymMember)
	if result.Error!=nil{
		fmt.Println("error in DB")

	}

	fmt.Fprint(w,"Gym Employ Enrolled!!!")

}

func Payment(w http.ResponseWriter,r *http.Request){

	params:=mux.Vars(r)

	var newPayment model.Payment

	json.NewDecoder(r.Body).Decode(&newPayment)

	
	newPayment.User_id=params["user_id"]
	newPayment.Date=time.Now()

	result:=db.Create(&newPayment)
	if result.Error!=nil{
		fmt.Println("error in DB")
	}

	fmt.Fprint(w,"Payment Successfull !!")

}

func ChooseSubscription(w http.ResponseWriter,r *http.Request){


	params:=mux.Vars(r)
	var subscription model.Subscription

	json.NewDecoder(r.Body).Decode(&subscription)

	subscription.User_id=params["user_id"]
	

	result:=db.Create(&subscription)
	if result.Error!=nil{
		fmt.Println("error in DB")
	}

	fmt.Fprint(w,"Subscription chosen !!")

}













//gorm DB global variable
var db *gorm.DB
var err error

// ------------------------------>(MAIN FUNCTION)
func main() {


	
router := mux.NewRouter()

dsn := "host=localhost port=5432 user=postgres password=6280912015 dbname=gorm_db sslmode=disable TimeZone=Asia/Shanghai"

db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})


if err != nil {
	panic("failed to connect database")
}

fmt.Println("DB connection established")

//defer db.Close()

db.AutoMigrate(&model.Equipment{},&model.Users{},&model.GymEmployee{},&model.Payment{},&model.Price{},&model.Subscription{})




// router.HandleFunc("/gym/enrollmentData",CustomerEnrolmentsData)
// router.HandleFunc("/gym/enrollmentData/{id}", CustomerEnrolmentsDatabyID)
router.HandleFunc("/gym/enrollment", CustomerEnrolments)
router.HandleFunc("/gym/GymEmployEnrolment",GymEmployEnrolment)
router.HandleFunc("/gym/enrollment/Subscription/payment/{user_id}",Payment)
router.HandleFunc("/gym/enrollment/Subscription/{user_id}",ChooseSubscription)

// router.HandleFunc("/gym/deleteMember/{id}", DeleteMembership)//
// router.HandleFunc("/gym/setPrice", SetMembershipPrice)
// router.HandleFunc("/gym/createMembershipPrice",CreateMembershipPriceDB)//post method




log.Fatal(http.ListenAndServe(":8888", router))
	// fmt.Println("dataBASE",dataBASE)



}