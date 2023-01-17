// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/blogs"
	"entdemo/ent/predicate"
	"entdemo/ent/user"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeBlogs = "Blogs"
	TypeUser  = "User"
)

// BlogsMutation represents an operation that mutates the Blogs nodes in the graph.
type BlogsMutation struct {
	config
	op            Op
	typ           string
	id            *int
	blogTitle     *string
	blogType      *string
	blogContent   *string
	blogAuthor    *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Blogs, error)
	predicates    []predicate.Blogs
}

var _ ent.Mutation = (*BlogsMutation)(nil)

// blogsOption allows management of the mutation configuration using functional options.
type blogsOption func(*BlogsMutation)

// newBlogsMutation creates new mutation for the Blogs entity.
func newBlogsMutation(c config, op Op, opts ...blogsOption) *BlogsMutation {
	m := &BlogsMutation{
		config:        c,
		op:            op,
		typ:           TypeBlogs,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withBlogsID sets the ID field of the mutation.
func withBlogsID(id int) blogsOption {
	return func(m *BlogsMutation) {
		var (
			err   error
			once  sync.Once
			value *Blogs
		)
		m.oldValue = func(ctx context.Context) (*Blogs, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Blogs.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withBlogs sets the old Blogs of the mutation.
func withBlogs(node *Blogs) blogsOption {
	return func(m *BlogsMutation) {
		m.oldValue = func(context.Context) (*Blogs, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m BlogsMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m BlogsMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *BlogsMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *BlogsMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Blogs.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetBlogTitle sets the "blogTitle" field.
func (m *BlogsMutation) SetBlogTitle(s string) {
	m.blogTitle = &s
}

// BlogTitle returns the value of the "blogTitle" field in the mutation.
func (m *BlogsMutation) BlogTitle() (r string, exists bool) {
	v := m.blogTitle
	if v == nil {
		return
	}
	return *v, true
}

// OldBlogTitle returns the old "blogTitle" field's value of the Blogs entity.
// If the Blogs object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BlogsMutation) OldBlogTitle(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBlogTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBlogTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBlogTitle: %w", err)
	}
	return oldValue.BlogTitle, nil
}

// ClearBlogTitle clears the value of the "blogTitle" field.
func (m *BlogsMutation) ClearBlogTitle() {
	m.blogTitle = nil
	m.clearedFields[blogs.FieldBlogTitle] = struct{}{}
}

// BlogTitleCleared returns if the "blogTitle" field was cleared in this mutation.
func (m *BlogsMutation) BlogTitleCleared() bool {
	_, ok := m.clearedFields[blogs.FieldBlogTitle]
	return ok
}

// ResetBlogTitle resets all changes to the "blogTitle" field.
func (m *BlogsMutation) ResetBlogTitle() {
	m.blogTitle = nil
	delete(m.clearedFields, blogs.FieldBlogTitle)
}

// SetBlogType sets the "blogType" field.
func (m *BlogsMutation) SetBlogType(s string) {
	m.blogType = &s
}

// BlogType returns the value of the "blogType" field in the mutation.
func (m *BlogsMutation) BlogType() (r string, exists bool) {
	v := m.blogType
	if v == nil {
		return
	}
	return *v, true
}

// OldBlogType returns the old "blogType" field's value of the Blogs entity.
// If the Blogs object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BlogsMutation) OldBlogType(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBlogType is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBlogType requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBlogType: %w", err)
	}
	return oldValue.BlogType, nil
}

// ClearBlogType clears the value of the "blogType" field.
func (m *BlogsMutation) ClearBlogType() {
	m.blogType = nil
	m.clearedFields[blogs.FieldBlogType] = struct{}{}
}

// BlogTypeCleared returns if the "blogType" field was cleared in this mutation.
func (m *BlogsMutation) BlogTypeCleared() bool {
	_, ok := m.clearedFields[blogs.FieldBlogType]
	return ok
}

// ResetBlogType resets all changes to the "blogType" field.
func (m *BlogsMutation) ResetBlogType() {
	m.blogType = nil
	delete(m.clearedFields, blogs.FieldBlogType)
}

// SetBlogContent sets the "blogContent" field.
func (m *BlogsMutation) SetBlogContent(s string) {
	m.blogContent = &s
}

// BlogContent returns the value of the "blogContent" field in the mutation.
func (m *BlogsMutation) BlogContent() (r string, exists bool) {
	v := m.blogContent
	if v == nil {
		return
	}
	return *v, true
}

// OldBlogContent returns the old "blogContent" field's value of the Blogs entity.
// If the Blogs object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BlogsMutation) OldBlogContent(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBlogContent is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBlogContent requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBlogContent: %w", err)
	}
	return oldValue.BlogContent, nil
}

// ClearBlogContent clears the value of the "blogContent" field.
func (m *BlogsMutation) ClearBlogContent() {
	m.blogContent = nil
	m.clearedFields[blogs.FieldBlogContent] = struct{}{}
}

// BlogContentCleared returns if the "blogContent" field was cleared in this mutation.
func (m *BlogsMutation) BlogContentCleared() bool {
	_, ok := m.clearedFields[blogs.FieldBlogContent]
	return ok
}

// ResetBlogContent resets all changes to the "blogContent" field.
func (m *BlogsMutation) ResetBlogContent() {
	m.blogContent = nil
	delete(m.clearedFields, blogs.FieldBlogContent)
}

// SetBlogAuthor sets the "blogAuthor" field.
func (m *BlogsMutation) SetBlogAuthor(s string) {
	m.blogAuthor = &s
}

// BlogAuthor returns the value of the "blogAuthor" field in the mutation.
func (m *BlogsMutation) BlogAuthor() (r string, exists bool) {
	v := m.blogAuthor
	if v == nil {
		return
	}
	return *v, true
}

// OldBlogAuthor returns the old "blogAuthor" field's value of the Blogs entity.
// If the Blogs object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BlogsMutation) OldBlogAuthor(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBlogAuthor is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBlogAuthor requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBlogAuthor: %w", err)
	}
	return oldValue.BlogAuthor, nil
}

// ClearBlogAuthor clears the value of the "blogAuthor" field.
func (m *BlogsMutation) ClearBlogAuthor() {
	m.blogAuthor = nil
	m.clearedFields[blogs.FieldBlogAuthor] = struct{}{}
}

// BlogAuthorCleared returns if the "blogAuthor" field was cleared in this mutation.
func (m *BlogsMutation) BlogAuthorCleared() bool {
	_, ok := m.clearedFields[blogs.FieldBlogAuthor]
	return ok
}

// ResetBlogAuthor resets all changes to the "blogAuthor" field.
func (m *BlogsMutation) ResetBlogAuthor() {
	m.blogAuthor = nil
	delete(m.clearedFields, blogs.FieldBlogAuthor)
}

// Where appends a list predicates to the BlogsMutation builder.
func (m *BlogsMutation) Where(ps ...predicate.Blogs) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the BlogsMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *BlogsMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Blogs, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *BlogsMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *BlogsMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Blogs).
func (m *BlogsMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *BlogsMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.blogTitle != nil {
		fields = append(fields, blogs.FieldBlogTitle)
	}
	if m.blogType != nil {
		fields = append(fields, blogs.FieldBlogType)
	}
	if m.blogContent != nil {
		fields = append(fields, blogs.FieldBlogContent)
	}
	if m.blogAuthor != nil {
		fields = append(fields, blogs.FieldBlogAuthor)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *BlogsMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case blogs.FieldBlogTitle:
		return m.BlogTitle()
	case blogs.FieldBlogType:
		return m.BlogType()
	case blogs.FieldBlogContent:
		return m.BlogContent()
	case blogs.FieldBlogAuthor:
		return m.BlogAuthor()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *BlogsMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case blogs.FieldBlogTitle:
		return m.OldBlogTitle(ctx)
	case blogs.FieldBlogType:
		return m.OldBlogType(ctx)
	case blogs.FieldBlogContent:
		return m.OldBlogContent(ctx)
	case blogs.FieldBlogAuthor:
		return m.OldBlogAuthor(ctx)
	}
	return nil, fmt.Errorf("unknown Blogs field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *BlogsMutation) SetField(name string, value ent.Value) error {
	switch name {
	case blogs.FieldBlogTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBlogTitle(v)
		return nil
	case blogs.FieldBlogType:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBlogType(v)
		return nil
	case blogs.FieldBlogContent:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBlogContent(v)
		return nil
	case blogs.FieldBlogAuthor:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBlogAuthor(v)
		return nil
	}
	return fmt.Errorf("unknown Blogs field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *BlogsMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *BlogsMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *BlogsMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Blogs numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *BlogsMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(blogs.FieldBlogTitle) {
		fields = append(fields, blogs.FieldBlogTitle)
	}
	if m.FieldCleared(blogs.FieldBlogType) {
		fields = append(fields, blogs.FieldBlogType)
	}
	if m.FieldCleared(blogs.FieldBlogContent) {
		fields = append(fields, blogs.FieldBlogContent)
	}
	if m.FieldCleared(blogs.FieldBlogAuthor) {
		fields = append(fields, blogs.FieldBlogAuthor)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *BlogsMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *BlogsMutation) ClearField(name string) error {
	switch name {
	case blogs.FieldBlogTitle:
		m.ClearBlogTitle()
		return nil
	case blogs.FieldBlogType:
		m.ClearBlogType()
		return nil
	case blogs.FieldBlogContent:
		m.ClearBlogContent()
		return nil
	case blogs.FieldBlogAuthor:
		m.ClearBlogAuthor()
		return nil
	}
	return fmt.Errorf("unknown Blogs nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *BlogsMutation) ResetField(name string) error {
	switch name {
	case blogs.FieldBlogTitle:
		m.ResetBlogTitle()
		return nil
	case blogs.FieldBlogType:
		m.ResetBlogType()
		return nil
	case blogs.FieldBlogContent:
		m.ResetBlogContent()
		return nil
	case blogs.FieldBlogAuthor:
		m.ResetBlogAuthor()
		return nil
	}
	return fmt.Errorf("unknown Blogs field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *BlogsMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *BlogsMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *BlogsMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *BlogsMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *BlogsMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *BlogsMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *BlogsMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Blogs unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *BlogsMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Blogs edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	password      *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// SetPassword sets the "password" field.
func (m *UserMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *UserMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword resets all changes to the "password" field.
func (m *UserMutation) ResetPassword() {
	m.password = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	if m.password != nil {
		fields = append(fields, user.FieldPassword)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldName:
		return m.Name()
	case user.FieldPassword:
		return m.Password()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldName:
		return m.OldName(ctx)
	case user.FieldPassword:
		return m.OldPassword(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case user.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldName:
		m.ResetName()
		return nil
	case user.FieldPassword:
		m.ResetPassword()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown User edge %s", name)
}
