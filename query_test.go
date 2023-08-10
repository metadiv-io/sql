package sql_test

import (
	"os"

	"github.com/google/uuid"
	"github.com/metadiv-io/sql"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("test query", func() {

	var DB *gorm.DB
	var pathName string

	type User struct {
		gorm.Model
		Name string
		Age  int
	}

	BeforeEach(func() {
		var err error
		pathName = uuid.New().String() + ".db"
		DB, err = sql.Connect(true).Sqlite(pathName)
		Expect(err).To(BeNil())

		err = DB.AutoMigrate(&User{})
		Expect(err).To(BeNil())

		err = DB.Create(&User{Name: "peter", Age: 18}).Error
		Expect(err).To(BeNil())

		err = DB.Create(&User{Name: "tom", Age: 20}).Error
		Expect(err).To(BeNil())

		err = DB.Create(&User{Name: "jerry", Age: 22}).Error
		Expect(err).To(BeNil())
	})

	AfterEach(func() {
		os.RemoveAll(pathName)
	})

	It("should be able to query all", func() {
		users, err := sql.FindAll[User](DB, nil)
		Expect(err).To(BeNil())
		Expect(len(users)).To(Equal(3))
		Expect(users[0].Name).To(Equal("peter"))
		Expect(users[1].Name).To(Equal("tom"))
		Expect(users[2].Name).To(Equal("jerry"))
	})

	It("should be able to query all with clause", func() {
		users, err := sql.FindAll[User](DB, sql.Eq("name", "peter"))
		Expect(err).To(BeNil())
		Expect(len(users)).To(Equal(1))
		Expect(users[0].Name).To(Equal("peter"))

		users, err = sql.FindAll[User](DB, sql.Gte("age", 20))
		Expect(err).To(BeNil())
		Expect(len(users)).To(Equal(2))
		Expect(users[0].Name).To(Equal("tom"))
		Expect(users[1].Name).To(Equal("jerry"))
	})

	It("should be able to query one with clause", func() {
		user, err := sql.FindOne[User](DB, sql.Eq("name", "peter"))
		Expect(err).To(BeNil())
		Expect(user.Name).To(Equal("peter"))
	})

	It("should be able to count", func() {
		count, err := sql.Count[User](DB, nil)
		Expect(err).To(BeNil())
		Expect(count).To(Equal(int64(3)))

		count, err = sql.Count[User](DB, sql.Eq("name", "peter"))
		Expect(err).To(BeNil())
		Expect(count).To(Equal(int64(1)))
	})

	It("should be able to query with order", func() {
		users, _, err := sql.FindAllComplex[User](DB, DB, nil, sql.Order("age", false), nil)
		Expect(err).To(BeNil())
		Expect(len(users)).To(Equal(3))
		Expect(users[0].Name).To(Equal("jerry"))
		Expect(users[1].Name).To(Equal("tom"))
		Expect(users[2].Name).To(Equal("peter"))
	})

	It("should be able to query with pagination", func() {
		users, pagination, err := sql.FindAllComplex[User](DB, DB, nil, nil, sql.Page(1, 2))
		Expect(err).To(BeNil())
		Expect(len(users)).To(Equal(2))
		Expect(users[0].Name).To(Equal("peter"))
		Expect(users[1].Name).To(Equal("tom"))
		Expect(pagination.Total).To(Equal(int64(3)))
		Expect(pagination.Page).To(Equal(1))
		Expect(pagination.Size).To(Equal(2))
	})

	It("should be able to query with order and pagination and clause", func() {
		users, pagination, err := sql.FindAllComplex[User](DB, DB, sql.Gte("age", 20), sql.Order("age", true), sql.Page(1, 2))
		Expect(err).To(BeNil())
		Expect(len(users)).To(Equal(2))
		Expect(users[0].Name).To(Equal("tom"))
		Expect(users[1].Name).To(Equal("jerry"))
		Expect(pagination.Total).To(Equal(int64(2)))
		Expect(pagination.Page).To(Equal(1))
		Expect(pagination.Size).To(Equal(2))
	})
})
