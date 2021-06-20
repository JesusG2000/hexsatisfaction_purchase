package service

import (
	"context"
	"testing"
	"time"

	"github.com/JesusG2000/hexsatisfaction_purchase/internal/model"
	m "github.com/JesusG2000/hexsatisfaction_purchase/internal/service/mock"
	"github.com/pkg/errors"
	testAssert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCommentService_Create(t *testing.T) {
	primitive.NewObjectID().Hex()
	assert := testAssert.New(t)
	type test struct {
		name   string
		req    model.CreateCommentRequest
		fn     func(comment *m.Comment, existanse *m.Existanse, data test)
		expID  string
		expErr error
	}
	tt := []test{
		{
			name: "check user error",
			req: model.CreateCommentRequest{
				UserID:     1,
				PurchaseID: primitive.NewObjectID().Hex(),
				Date:       time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				Text:       "some text",
			},
			fn: func(comment *m.Comment, existanse *m.Existanse, data test) {
				existanse.On("User", mock.Anything, data.req.UserID).
					Return(false, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "check error"),
		},
		{
			name: "user  found",
			req: model.CreateCommentRequest{
				UserID:     1,
				PurchaseID: primitive.NewObjectID().Hex(),
				Date:       time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				Text:       "some text",
			},
			fn: func(comment *m.Comment, existanse *m.Existanse, data test) {
				existanse.On("User", mock.Anything, data.req.UserID).
					Return(true, nil)
			},
		},
		{
			name: "create error",
			req: model.CreateCommentRequest{
				UserID:     1,
				PurchaseID: primitive.NewObjectID().Hex(),
				Date:       time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				Text:       "some text",
			},
			fn: func(comment *m.Comment, existanse *m.Existanse, data test) {
				existanse.On("User", mock.Anything, data.req.UserID).
					Return(false, nil)
				comment.On("Create", mock.Anything, model.CommentDTO{
					UserID:     data.req.UserID,
					PurchaseID: data.req.PurchaseID,
					Date:       data.req.Date,
					Text:       data.req.Text,
				}).
					Return("", errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't create comment"),
		},
		{
			name: "All ok",
			req: model.CreateCommentRequest{
				UserID:     1,
				PurchaseID: primitive.NewObjectID().Hex(),
				Date:       time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				Text:       "some text",
			},
			fn: func(comment *m.Comment, existanse *m.Existanse, data test) {
				existanse.On("User", mock.Anything, data.req.UserID).
					Return(false, nil)
				comment.On("Create", mock.Anything, model.CommentDTO{
					UserID:     data.req.UserID,
					PurchaseID: data.req.PurchaseID,
					Date:       data.req.Date,
					Text:       data.req.Text,
				}).
					Return(data.expID, nil)
			},
			expID: primitive.NewObjectID().Hex(),
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			comment := new(m.Comment)
			existanse := new(m.Existanse)
			ctx := context.Background()
			service := NewCommentService(comment, existanse)
			if tc.fn != nil {
				tc.fn(comment, existanse, tc)
			}
			id, err := service.Create(ctx, tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.expID, id)
		})
	}
}

func TestCommentService_Update(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name   string
		req    model.UpdateCommentRequest
		fn     func(comment *m.Comment, existanse *m.Existanse, data test)
		expID  string
		expErr error
	}
	tt := []test{
		{
			name: "check user errors",
			req: model.UpdateCommentRequest{
				ID:         primitive.NewObjectID().Hex(),
				UserID:     1,
				PurchaseID: primitive.NewObjectID().Hex(),
				Date:       time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				Text:       "some text",
			},
			fn: func(comment *m.Comment, existanse *m.Existanse, data test) {
				existanse.On("User", mock.Anything, data.req.UserID).
					Return(false, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "check error"),
		},
		{
			name: "user not exist",
			req: model.UpdateCommentRequest{
				ID:         primitive.NewObjectID().Hex(),
				UserID:     1,
				PurchaseID: primitive.NewObjectID().Hex(),
				Date:       time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				Text:       "some text",
			},
			fn: func(comment *m.Comment, existanse *m.Existanse, data test) {
				existanse.On("User", mock.Anything, data.req.UserID).
					Return(false, nil)
			},
		},
		{
			name: "Update errors",
			req: model.UpdateCommentRequest{
				ID:         primitive.NewObjectID().Hex(),
				UserID:     1,
				PurchaseID: primitive.NewObjectID().Hex(),
				Date:       time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				Text:       "some text",
			},
			fn: func(comment *m.Comment, existanse *m.Existanse, data test) {
				existanse.On("User", mock.Anything, data.req.UserID).
					Return(true, nil)
				comment.On("Update", mock.Anything, data.req.ID, model.CommentDTO{
					UserID:     data.req.UserID,
					PurchaseID: data.req.PurchaseID,
					Date:       data.req.Date,
					Text:       data.req.Text,
				}).
					Return(data.expID, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't update comment"),
		},
		{
			name: "All ok",
			req: model.UpdateCommentRequest{
				ID:         primitive.NewObjectID().Hex(),
				UserID:     1,
				PurchaseID: primitive.NewObjectID().Hex(),
				Date:       time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				Text:       "some text",
			},
			fn: func(comment *m.Comment, existanse *m.Existanse, data test) {
				existanse.On("User", mock.Anything, data.req.UserID).
					Return(true, nil)
				comment.On("Update", mock.Anything, data.req.ID, model.CommentDTO{
					UserID:     data.req.UserID,
					PurchaseID: data.req.PurchaseID,
					Date:       data.req.Date,
					Text:       data.req.Text,
				}).
					Return(data.expID, nil)
			},
			expID: primitive.NewObjectID().Hex(),
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			comment := new(m.Comment)
			existence := new(m.Existanse)
			ctx := context.Background()
			service := NewCommentService(comment, existence)
			if tc.fn != nil {
				tc.fn(comment, existence, tc)
			}
			id, err := service.Update(ctx, tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.expID, id)
		})
	}
}

func TestCommentService_Delete(t *testing.T) {
	assert := testAssert.New(t)

	type test struct {
		name   string
		req    model.DeleteCommentRequest
		fn     func(comment *m.Comment, data test)
		expID  string
		expErr error
	}
	tt := []test{
		{
			name: "Delete errors",
			req: model.DeleteCommentRequest{
				ID: primitive.NewObjectID().Hex(),
			},
			fn: func(comment *m.Comment, data test) {
				comment.On("Delete", mock.Anything, data.req.ID).
					Return(data.expID, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't delete comment"),
		},
		{
			name: "All ok",
			req: model.DeleteCommentRequest{
				ID: primitive.NewObjectID().Hex(),
			},
			fn: func(comment *m.Comment, data test) {
				comment.On("Delete", mock.Anything, data.req.ID).
					Return(data.expID, nil)
			},
			expID: primitive.NewObjectID().Hex(),
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			comment := new(m.Comment)
			existence := new(m.Existanse)
			ctx := context.Background()
			service := NewCommentService(comment, existence)
			if tc.fn != nil {
				tc.fn(comment, tc)
			}
			id, err := service.Delete(ctx, tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.expID, id)
		})
	}
}

func TestCommentService_FindByID(t *testing.T) {
	assert := testAssert.New(t)

	type test struct {
		name   string
		req    model.IDCommentRequest
		fn     func(comment *m.Comment, data test)
		exp    *model.CommentDTO
		expErr error
	}
	tt := []test{
		{
			name: "Find errors",
			req: model.IDCommentRequest{
				ID: primitive.NewObjectID().Hex(),
			},
			fn: func(comment *m.Comment, data test) {
				comment.On("FindByID", mock.Anything, data.req.ID).
					Return(data.exp, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find comment"),
		},
		{
			name: "All ok",
			req: model.IDCommentRequest{
				ID: primitive.NewObjectID().Hex(),
			},
			fn: func(comment *m.Comment, data test) {
				comment.On("FindByID", mock.Anything, data.req.ID).
					Return(data.exp, nil)
			},
			exp: &model.CommentDTO{
				ID:         primitive.NewObjectID().Hex(),
				UserID:     1,
				PurchaseID: primitive.NewObjectID().Hex(),
				Date:       time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
				Text:       "some text",
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			comment := new(m.Comment)
			existence := new(m.Existanse)
			ctx := context.Background()
			service := NewCommentService(comment, existence)
			if tc.fn != nil {
				tc.fn(comment, tc)
			}
			c, err := service.FindByID(ctx, tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.exp, c)
		})
	}
}

func TestCommentService_FindAllByUserID(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name   string
		req    model.UserIDCommentRequest
		fn     func(comment *m.Comment, existanse *m.Existanse, data *test)
		exp    []model.CommentDTO
		expErr error
	}
	tt := []test{
		{
			name: "check error",
			req: model.UserIDCommentRequest{
				ID: 1,
			},
			fn: func(comment *m.Comment, existanse *m.Existanse, data *test) {
				existanse.On("User", mock.Anything, data.req.ID).
					Return(false, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "check error"),
		},
		{
			name: "user not found",
			req: model.UserIDCommentRequest{
				ID: 1,
			},
			fn: func(comment *m.Comment, existanse *m.Existanse, data *test) {
				existanse.On("User", mock.Anything, data.req.ID).
					Return(false, nil)
			},
		},
		{
			name: "Find errors",
			req: model.UserIDCommentRequest{
				ID: 1,
			},
			fn: func(comment *m.Comment, existanse *m.Existanse, data *test) {
				existanse.On("User", mock.Anything, data.req.ID).
					Return(true, nil)
				comment.On("FindAllByUserID", mock.Anything, data.req.ID).
					Return(data.exp, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find comments"),
		},
		{
			name: "All ok",
			req: model.UserIDCommentRequest{
				ID: 1,
			},
			fn: func(comment *m.Comment, existanse *m.Existanse, data *test) {
				for i := range data.exp {
					data.exp[i].UserID = data.req.ID
				}
				existanse.On("User", mock.Anything, data.req.ID).
					Return(true, nil)
				comment.On("FindAllByUserID", mock.Anything, data.req.ID).
					Return(data.exp, nil)
			},
			exp: []model.CommentDTO{
				{
					ID:         primitive.NewObjectID().Hex(),
					PurchaseID: primitive.NewObjectID().Hex(),
					Date:       time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					Text:       "some text1",
				},
				{
					ID:         primitive.NewObjectID().Hex(),
					PurchaseID: primitive.NewObjectID().Hex(),
					Date:       time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					Text:       "some text2",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			comment := new(m.Comment)
			existence := new(m.Existanse)
			ctx := context.Background()
			service := NewCommentService(comment, existence)
			if tc.fn != nil {
				tc.fn(comment, existence, &tc)
			}
			c, err := service.FindAllByUserID(ctx, tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.exp, c)
		})
	}
}

func TestCommentService_FindAllByPurchaseID(t *testing.T) {
	assert := testAssert.New(t)

	type test struct {
		name   string
		req    model.PurchaseIDCommentRequest
		fn     func(comment *m.Comment, data *test)
		exp    []model.CommentDTO
		expErr error
	}
	tt := []test{
		{
			name: "Find errors",
			req: model.PurchaseIDCommentRequest{
				ID: primitive.NewObjectID().Hex(),
			},
			fn: func(comment *m.Comment, data *test) {
				comment.On("FindByPurchaseID", mock.Anything, data.req.ID).
					Return(data.exp, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find comments"),
		},
		{
			name: "All ok",
			req: model.PurchaseIDCommentRequest{
				ID: primitive.NewObjectID().Hex(),
			},
			fn: func(comment *m.Comment, data *test) {
				for i := range data.exp {
					data.exp[i].PurchaseID = data.req.ID
				}
				comment.On("FindByPurchaseID", mock.Anything, data.req.ID).
					Return(data.exp, nil)
			},
			exp: []model.CommentDTO{
				{
					ID:     primitive.NewObjectID().Hex(),
					UserID: 1,
					Date:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					Text:   "some text1",
				},
				{
					ID:     primitive.NewObjectID().Hex(),
					UserID: 1,
					Date:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					Text:   "some text2",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			comment := new(m.Comment)
			existence := new(m.Existanse)
			ctx := context.Background()
			service := NewCommentService(comment, existence)
			if tc.fn != nil {
				tc.fn(comment, &tc)
			}
			c, err := service.FindByPurchaseID(ctx, tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.exp, c)
		})
	}
}

func TestCommentService_FindByUserIDAndPurchaseID(t *testing.T) {
	assert := testAssert.New(t)
	type test struct {
		name   string
		req    model.UserPurchaseIDCommentRequest
		fn     func(comment *m.Comment, existanse *m.Existanse, data *test)
		exp    []model.CommentDTO
		expErr error
	}
	tt := []test{
		{
			name: "check error",
			req: model.UserPurchaseIDCommentRequest{
				UserID:     1,
				PurchaseID: primitive.NewObjectID().Hex(),
			},
			fn: func(comment *m.Comment, existanse *m.Existanse, data *test) {
				existanse.On("User", mock.Anything, data.req.UserID).
					Return(false, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "check error"),
		},
		{
			name: "user not found",
			req: model.UserPurchaseIDCommentRequest{
				UserID:     1,
				PurchaseID: primitive.NewObjectID().Hex(),
			},
			fn: func(comment *m.Comment, existanse *m.Existanse, data *test) {
				existanse.On("User", mock.Anything, data.req.UserID).
					Return(false, nil)
			},
		},
		{
			name: "Find errors",
			req: model.UserPurchaseIDCommentRequest{
				UserID:     1,
				PurchaseID: primitive.NewObjectID().Hex(),
			},
			fn: func(comment *m.Comment, existanse *m.Existanse, data *test) {
				existanse.On("User", mock.Anything, data.req.UserID).
					Return(true, nil)
				comment.On("FindByUserIDAndPurchaseID", mock.Anything, data.req.UserID, data.req.PurchaseID).
					Return(data.exp, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find comments"),
		},
		{
			name: "All ok",
			req: model.UserPurchaseIDCommentRequest{
				UserID:     1,
				PurchaseID: primitive.NewObjectID().Hex(),
			},
			fn: func(comment *m.Comment, existanse *m.Existanse, data *test) {
				for i := range data.exp {
					data.exp[i].UserID = data.req.UserID
					data.exp[i].PurchaseID = data.req.PurchaseID
				}
				existanse.On("User", mock.Anything, data.req.UserID).
					Return(true, nil)
				comment.On("FindByUserIDAndPurchaseID", mock.Anything, data.req.UserID, data.req.PurchaseID).
					Return(data.exp, nil)
			},
			exp: []model.CommentDTO{
				{
					ID:   primitive.NewObjectID().Hex(),
					Date: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					Text: "some text1",
				},
				{
					ID:   primitive.NewObjectID().Hex(),
					Date: time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					Text: "some text2",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			comment := new(m.Comment)
			existence := new(m.Existanse)
			ctx := context.Background()
			service := NewCommentService(comment, existence)
			if tc.fn != nil {
				tc.fn(comment, existence, &tc)
			}
			c, err := service.FindByUserIDAndPurchaseID(ctx, tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.exp, c)
		})
	}
}

func TestCommentService_FindAll(t *testing.T) {
	assert := testAssert.New(t)

	type test struct {
		name   string
		fn     func(comment *m.Comment, data test)
		exp    []model.CommentDTO
		expErr error
	}
	tt := []test{
		{
			name: "Find errors",

			fn: func(comment *m.Comment, data test) {
				comment.On("FindAll", mock.Anything).
					Return(data.exp, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find comments"),
		},
		{
			name: "All ok",
			fn: func(comment *m.Comment, data test) {
				comment.On("FindAll", mock.Anything).
					Return(data.exp, nil)
			},
			exp: []model.CommentDTO{
				{
					ID:         primitive.NewObjectID().Hex(),
					UserID:     1,
					PurchaseID: primitive.NewObjectID().Hex(),
					Date:       time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					Text:       "some text1",
				},
				{
					ID:         primitive.NewObjectID().Hex(),
					UserID:     1,
					PurchaseID: primitive.NewObjectID().Hex(),
					Date:       time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					Text:       "some text2",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			comment := new(m.Comment)
			existence := new(m.Existanse)
			ctx := context.Background()
			service := NewCommentService(comment, existence)
			if tc.fn != nil {
				tc.fn(comment, tc)
			}
			c, err := service.FindAll(ctx)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.exp, c)
		})
	}
}

func TestCommentService_FindByText(t *testing.T) {
	assert := testAssert.New(t)

	type test struct {
		name   string
		req    model.TextCommentRequest
		fn     func(comment *m.Comment, data test)
		exp    []model.CommentDTO
		expErr error
	}
	tt := []test{
		{
			name: "Find errors",
			req: model.TextCommentRequest{
				Text: "some",
			},
			fn: func(comment *m.Comment, data test) {
				comment.On("FindByText", mock.Anything, data.req.Text).
					Return(data.exp, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find comments"),
		},
		{
			name: "All ok",
			req: model.TextCommentRequest{
				Text: "some",
			},
			fn: func(comment *m.Comment, data test) {
				comment.On("FindByText", mock.Anything, data.req.Text).
					Return(data.exp, nil)
			},
			exp: []model.CommentDTO{
				{
					ID:         primitive.NewObjectID().Hex(),
					UserID:     1,
					PurchaseID: primitive.NewObjectID().Hex(),
					Date:       time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					Text:       "some text1",
				},
				{
					ID:         primitive.NewObjectID().Hex(),
					UserID:     1,
					PurchaseID: primitive.NewObjectID().Hex(),
					Date:       time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					Text:       "some text2",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			comment := new(m.Comment)
			existence := new(m.Existanse)
			ctx := context.Background()
			service := NewCommentService(comment, existence)
			if tc.fn != nil {
				tc.fn(comment, tc)
			}
			c, err := service.FindByText(ctx, tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.exp, c)
		})
	}
}

func TestCommentService_FindByPeriod(t *testing.T) {
	assert := testAssert.New(t)

	type test struct {
		name   string
		req    model.PeriodCommentRequest
		fn     func(comment *m.Comment, data test)
		exp    []model.CommentDTO
		expErr error
	}
	tt := []test{
		{
			name: "Find errors",
			req: model.PeriodCommentRequest{
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				End:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(comment *m.Comment, data test) {
				comment.On("FindByPeriod", mock.Anything, data.req.Start, data.req.End).
					Return(data.exp, errors.New(""))
			},
			expErr: errors.Wrap(errors.New(""), "couldn't find comments"),
		},
		{
			name: "All ok",
			req: model.PeriodCommentRequest{
				Start: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
				End:   time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
			},
			fn: func(comment *m.Comment, data test) {
				comment.On("FindByPeriod", mock.Anything, data.req.Start, data.req.End).
					Return(data.exp, nil)
			},
			exp: []model.CommentDTO{
				{
					ID:         primitive.NewObjectID().Hex(),
					UserID:     1,
					PurchaseID: primitive.NewObjectID().Hex(),
					Date:       time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local),
					Text:       "some text1",
				},
				{
					ID:         primitive.NewObjectID().Hex(),
					UserID:     1,
					PurchaseID: primitive.NewObjectID().Hex(),
					Date:       time.Date(2009, time.December, 10, 23, 0, 0, 0, time.Local),
					Text:       "some text2",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			comment := new(m.Comment)
			existence := new(m.Existanse)
			ctx := context.Background()
			service := NewCommentService(comment, existence)
			if tc.fn != nil {
				tc.fn(comment, tc)
			}
			c, err := service.FindByPeriod(ctx, tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}
			assert.Equal(tc.exp, c)
		})
	}
}
