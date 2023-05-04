package sql_test

import (
	"github.com/metadiv-io/sql"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("test clause builder", func() {

	It("should be able to build a clause (=)", func() {
		clause := sql.Eq("field", "value")
		stm, values := clause.Build()
		Expect(stm).To(Equal("`field` = ?"))
		Expect(values).To(Equal([]interface{}{"value"}))
	})

	It("should be able to build a clause (<>)", func() {
		clause := sql.Neq("field", "value")
		stm, values := clause.Build()
		Expect(stm).To(Equal("`field` <> ?"))
		Expect(values).To(Equal([]interface{}{"value"}))
	})

	It("should be able to build a clause (>)", func() {
		clause := sql.Gt("field", "value")
		stm, values := clause.Build()
		Expect(stm).To(Equal("`field` > ?"))
		Expect(values).To(Equal([]interface{}{"value"}))
	})

	It("should be able to build a clause (<)", func() {
		clause := sql.Lt("field", "value")
		stm, values := clause.Build()
		Expect(stm).To(Equal("`field` < ?"))
		Expect(values).To(Equal([]interface{}{"value"}))
	})

	It("should be able to build a clause (>=)", func() {
		clause := sql.Gte("field", "value")
		stm, values := clause.Build()
		Expect(stm).To(Equal("`field` >= ?"))
		Expect(values).To(Equal([]interface{}{"value"}))
	})

	It("should be able to build a clause (<=)", func() {
		clause := sql.Lte("field", "value")
		stm, values := clause.Build()
		Expect(stm).To(Equal("`field` <= ?"))
		Expect(values).To(Equal([]interface{}{"value"}))
	})

	It("should be able to build a clause (LIKE)", func() {
		clause := sql.Like("field", "value")
		stm, values := clause.Build()
		Expect(stm).To(Equal("`field` LIKE ?"))
		Expect(values).To(Equal([]interface{}{"value"}))
	})

	It("should be able to build a clause (NOT LIKE)", func() {
		clause := sql.NotLike("field", "value")
		stm, values := clause.Build()
		Expect(stm).To(Equal("`field` NOT LIKE ?"))
		Expect(values).To(Equal([]interface{}{"value"}))
	})

	It("should be able to build a clause (LIKE %%)", func() {
		clause := sql.Similar("field", "value")
		stm, values := clause.Build()
		Expect(stm).To(Equal("`field` LIKE ?"))
		Expect(values).To(Equal([]interface{}{"%value%"}))
	})

	It("should be able to build a clause (NOT LIKE %%)", func() {
		clause := sql.NotSimilar("field", "value")
		stm, values := clause.Build()
		Expect(stm).To(Equal("`field` NOT LIKE ?"))
		Expect(values).To(Equal([]interface{}{"%value%"}))
	})

	It("should be able to build a clause (IN)", func() {
		clause := sql.In("field", "value1", "value2")
		stm, values := clause.Build()
		Expect(stm).To(Equal("`field` IN (?,?)"))
		Expect(values).To(Equal([]interface{}{"value1", "value2"}))
	})

	It("should be able to build a clause (NOT IN)", func() {
		clause := sql.NotIn("field", "value1", "value2")
		stm, values := clause.Build()
		Expect(stm).To(Equal("`field` NOT IN (?,?)"))
		Expect(values).To(Equal([]interface{}{"value1", "value2"}))
	})

	It("should be able to build a clause (IS NULL)", func() {
		clause := sql.IsNull("field")
		stm, values := clause.Build()
		Expect(stm).To(Equal("`field` IS NULL"))
		Expect(len(values)).To(Equal(0))
	})

	It("should be able to build a clause (IS NOT NULL)", func() {
		clause := sql.IsNotNull("field")
		stm, values := clause.Build()
		Expect(stm).To(Equal("`field` IS NOT NULL"))
		Expect(len(values)).To(Equal(0))
	})

	It("should be able to build a clause (AND)", func() {
		clause := sql.And(
			sql.Eq("field1", "value1"),
			sql.Eq("field2", "value2"),
		)
		stm, values := clause.Build()
		Expect(stm).To(Equal("(`field1` = ? AND `field2` = ?)"))
		Expect(values).To(Equal([]interface{}{"value1", "value2"}))
	})

	It("should be able to build a clause (OR)", func() {
		clause := sql.Or(
			sql.Eq("field1", "value1"),
			sql.Eq("field2", "value2"),
		)
		stm, values := clause.Build()
		Expect(stm).To(Equal("(`field1` = ? OR `field2` = ?)"))
		Expect(values).To(Equal([]interface{}{"value1", "value2"}))
	})
})
