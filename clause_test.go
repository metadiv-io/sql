package sql_test

import (
	"github.com/metadiv-io/sql"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("test clause", func() {

	It("should be able to build a clause (=)", func() {
		clause := sql.NewClause("field", sql.OPERATOR_EQUAL, "value")
		stm, values := clause.Build()
		Expect(stm).To(Equal("`field` = ?"))
		Expect(values).To(Equal([]interface{}{"value"}))
	})

	It("should be able to build a clause (<>)", func() {
		clause := sql.NewClause("field", sql.OPERATOR_NOT_EQUAL, "value")
		stm, values := clause.Build()
		Expect(stm).To(Equal("`field` <> ?"))
		Expect(values).To(Equal([]interface{}{"value"}))
	})

	It("should be able to build a clause (>)", func() {
		clause := sql.NewClause("field", sql.OPERATOR_GREATER, "value")
		stm, values := clause.Build()
		Expect(stm).To(Equal("`field` > ?"))
		Expect(values).To(Equal([]interface{}{"value"}))
	})

	It("should be able to build a clause (<)", func() {
		clause := sql.NewClause("field", sql.OPERATOR_LESS, "value")
		stm, values := clause.Build()
		Expect(stm).To(Equal("`field` < ?"))
		Expect(values).To(Equal([]interface{}{"value"}))
	})

	It("should be able to build a clause (>=)", func() {
		clause := sql.NewClause("field", sql.OPERATOR_GREATER_EQUAL, "value")
		stm, values := clause.Build()
		Expect(stm).To(Equal("`field` >= ?"))
		Expect(values).To(Equal([]interface{}{"value"}))
	})

	It("should be able to build a clause (<=)", func() {
		clause := sql.NewClause("field", sql.OPERATOR_LESS_EQUAL, "value")
		stm, values := clause.Build()
		Expect(stm).To(Equal("`field` <= ?"))
		Expect(values).To(Equal([]interface{}{"value"}))
	})

	It("should be able to build a clause (LIKE)", func() {
		clause := sql.NewClause("field", sql.OPERATOR_LIKE, "value")
		stm, values := clause.Build()
		Expect(stm).To(Equal("`field` LIKE ?"))
		Expect(values).To(Equal([]interface{}{"value"}))
	})

	It("should be able to build a clause (NOT LIKE)", func() {
		clause := sql.NewClause("field", sql.OPERATOR_NOT_LIKE, "value")
		stm, values := clause.Build()
		Expect(stm).To(Equal("`field` NOT LIKE ?"))
		Expect(values).To(Equal([]interface{}{"value"}))
	})

	It("should be able to build a clause (IN)", func() {
		clause := sql.NewClause("field", sql.OPERATOR_IN, []interface{}{"value1", "value2"})
		stm, values := clause.Build()
		Expect(stm).To(Equal("`field` IN (?,?)"))
		Expect(values).To(Equal([]interface{}{"value1", "value2"}))
	})

	It("should be able to build a clause (NOT IN)", func() {
		clause := sql.NewClause("field", sql.OPERATOR_NOT_IN, []interface{}{"value1", "value2"})
		stm, values := clause.Build()
		Expect(stm).To(Equal("`field` NOT IN (?,?)"))
		Expect(values).To(Equal([]interface{}{"value1", "value2"}))
	})

	It("should be able to build a clause (IS NULL)", func() {
		clause := sql.NewClause("field", sql.OPERATOR_IS_NULL, nil)
		stm, values := clause.Build()
		Expect(stm).To(Equal("`field` IS NULL"))
		Expect(values).To(Equal([]interface{}{}))
	})

	It("should be able to build a clause (IS NOT NULL)", func() {
		clause := sql.NewClause("field", sql.OPERATOR_IS_NOT_NULL, nil)
		stm, values := clause.Build()
		Expect(stm).To(Equal("`field` IS NOT NULL"))
		Expect(values).To(Equal([]interface{}{}))
	})

	It("should be able to build a clause (AND)", func() {
		clause := sql.NewClause("field", sql.OPERATOR_AND, nil,
			sql.NewClause("field1", sql.OPERATOR_EQUAL, "value1"),
			sql.NewClause("field2", sql.OPERATOR_EQUAL, "value2"),
		)
		stm, values := clause.Build()
		Expect(stm).To(Equal("(`field1` = ? AND `field2` = ?)"))
		Expect(values).To(Equal([]interface{}{"value1", "value2"}))
	})

	It("should be able to build a clause (OR)", func() {
		clause := sql.NewClause("field", sql.OPERATOR_OR, nil,
			sql.NewClause("field1", sql.OPERATOR_EQUAL, "value1"),
			sql.NewClause("field2", sql.OPERATOR_EQUAL, "value2"),
		)
		stm, values := clause.Build()
		Expect(stm).To(Equal("(`field1` = ? OR `field2` = ?)"))
		Expect(values).To(Equal([]interface{}{"value1", "value2"}))
	})
})
