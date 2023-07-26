package sql_test

import (
	"os"

	"github.com/google/uuid"
	"github.com/metadiv-io/sql"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("test write", func() {

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

	It("should be able to create", func() {
		var err error
		user := &User{Name: "jack", Age: 18}
		user, err = sql.Save(DB, user)
		Expect(err).To(BeNil())
		Expect(user.ID).ToNot(BeZero())

		user, err = sql.FindOne[User](DB, sql.Eq("name", "jack"))
		Expect(err).To(BeNil())
		Expect(user.Name).To(Equal("jack"))
		Expect(user.Age).To(Equal(18))
		Expect(user.ID).ToNot(BeZero())
	})

	It("should be able to create all", func() {
		var err error
		users := []User{
			{Name: "jack", Age: 18},
			{Name: "rose", Age: 20},
		}
		users, err = sql.SaveAll(DB, users)
		Expect(err).To(BeNil())
		Expect(len(users)).To(Equal(2))
		Expect(users[0].ID).ToNot(BeZero())
		Expect(users[1].ID).ToNot(BeZero())

		users, err = sql.FindAll[User](DB, nil)
		Expect(err).To(BeNil())
		Expect(len(users)).To(Equal(5))
		Expect(users[0].Name).To(Equal("peter"))
		Expect(users[1].Name).To(Equal("tom"))
		Expect(users[2].Name).To(Equal("jerry"))
		Expect(users[3].Name).To(Equal("jack"))
		Expect(users[4].Name).To(Equal("rose"))
	})

	It("should be able to update", func() {
		user, err := sql.FindOne[User](DB, sql.Eq("name", "peter"))
		Expect(err).To(BeNil())
		Expect(user.Name).To(Equal("peter"))
		Expect(user.Age).To(Equal(18))
		Expect(user.ID).ToNot(BeZero())

		var id uint = user.ID

		user.Name = "jack"
		user.Age = 20
		user, err = sql.Save(DB, user)
		Expect(err).To(BeNil())
		Expect(user.Name).To(Equal("jack"))
		Expect(user.Age).To(Equal(20))
		Expect(user.ID).To(Equal(id))

		user, err = sql.FindOne[User](DB, sql.Eq("name", "jack"))
		Expect(err).To(BeNil())
		Expect(user.Name).To(Equal("jack"))
		Expect(user.Age).To(Equal(20))
		Expect(user.ID).To(Equal(id))
	})

	It("should be able to update all", func() {
		users, err := sql.FindAll[User](DB, sql.Lte("age", 20))
		Expect(err).To(BeNil())
		Expect(len(users)).To(Equal(2))

		users[0].Name = "jack"
		users[1].Name = "rose"
		users, err = sql.SaveAll(DB, users)
		Expect(err).To(BeNil())
		Expect(len(users)).To(Equal(2))

		users, err = sql.FindAll[User](DB, sql.Lte("age", 20))
		Expect(err).To(BeNil())
		Expect(len(users)).To(Equal(2))
		Expect(users[0].Name).To(Equal("jack"))
		Expect(users[1].Name).To(Equal("rose"))
	})

	It("should be able to delete", func() {
		user, err := sql.FindOne[User](DB, sql.Eq("name", "peter"))
		Expect(err).To(BeNil())
		Expect(user.Name).To(Equal("peter"))
		Expect(user.Age).To(Equal(18))
		Expect(user.ID).ToNot(BeZero())

		err = sql.Delete(DB, user)
		Expect(err).To(BeNil())

		_, err = sql.FindOne[User](DB, sql.Eq("name", "peter"))
		Expect(err).NotTo(BeNil())
	})

	It("should be able to delete all", func() {
		users, err := sql.FindAll[User](DB, sql.Lte("age", 20))
		Expect(err).To(BeNil())
		Expect(len(users)).To(Equal(2))

		err = sql.DeleteAll(DB, users)
		Expect(err).To(BeNil())

		count, err := sql.Count[User](DB, nil)
		Expect(err).To(BeNil())
		Expect(count).To(Equal(int64(1)))
	})

	It("should be able to delete with clause", func() {
		err := sql.DeleteAllByClause[User](DB, sql.Gte("age", 20))
		Expect(err).To(BeNil())

		count, err := sql.Count[User](DB, nil)
		Expect(err).To(BeNil())
		Expect(count).To(Equal(int64(1)))
	})
})
