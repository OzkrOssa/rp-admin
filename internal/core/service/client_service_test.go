package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/OzkrOssa/rp-admin/internal/core/domain"
	"github.com/OzkrOssa/rp-admin/internal/core/port/mock"
	"github.com/OzkrOssa/rp-admin/internal/core/service"
	"github.com/OzkrOssa/rp-admin/internal/core/util"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type registerTestedInput struct {
	client *domain.Client
}

type registerExpectedOutput struct {
	client *domain.Client
	err    error
}

func TestClientService_Register(t *testing.T) {
	ctx := context.Background()

	c := &domain.Client{
		FirstName:        gofakeit.Person().FirstName,
		LastName:         &gofakeit.Person().LastName,
		ClientTypeID:     uint64(1),
		Document:         gofakeit.SSN(),
		Email:            gofakeit.Email(),
		Profession:       &gofakeit.Person().Job.Title,
		Address:          gofakeit.Address().Address,
		Precinct:         gofakeit.Address().Street,
		MunicipalityCode: 20,
		DepartmentCode:   30,
		StatusID:         1,
		PhoneNumbers: []string{
			gofakeit.Phone(),
		},
	}

	clientInput := &domain.Client{
		FirstName:        c.FirstName,
		LastName:         c.LastName,
		ClientTypeID:     c.ClientTypeID,
		Document:         c.Document,
		Email:            c.Email,
		Profession:       c.Profession,
		MunicipalityCode: c.MunicipalityCode,
		DepartmentCode:   c.DepartmentCode,
		Address:          c.Address,
		Precinct:         c.Precinct,
		StatusID:         c.StatusID,
		PhoneNumbers:     c.PhoneNumbers,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	clientOutput := &domain.Client{
		ID:               gofakeit.Uint64(),
		FirstName:        c.FirstName,
		LastName:         c.LastName,
		ClientTypeID:     1,
		Document:         c.Document,
		Email:            c.Email,
		Profession:       c.Profession,
		MunicipalityCode: c.MunicipalityCode,
		DepartmentCode:   c.DepartmentCode,
		Address:          c.Address,
		Precinct:         c.Precinct,
		StatusID:         c.StatusID,
		PhoneNumbers:     c.PhoneNumbers,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	cacheKey := util.GenerateCacheKey("client", clientOutput.ID)
	userSerialized, _ := util.Serialize(clientOutput)
	ttl := time.Duration(0)

	testCases := []struct {
		desc  string
		mocks func(
			clientRepo *mock.MockClientRepository,
			cache *mock.MockCacheRepository,
		)
		input    registerTestedInput
		expected registerExpectedOutput
	}{
		{
			desc: "Success",
			mocks: func(clientRepo *mock.MockClientRepository, cache *mock.MockCacheRepository) {
				clientRepo.EXPECT().CreateClient(gomock.Any(), gomock.Eq(clientInput)).Return(clientOutput, nil)
				cache.EXPECT().Set(gomock.Any(), gomock.Eq(cacheKey), gomock.Eq(userSerialized), gomock.Eq(ttl)).Return(nil)
				cache.EXPECT().DeleteByPrefix(gomock.Any(), gomock.Eq("clients:*")).Return(nil)
			},
			input: registerTestedInput{
				clientInput,
			},
			expected: registerExpectedOutput{
				client: clientOutput,
				err:    nil,
			},
		},
		{
			desc: "Fail_DuplicateData",
			mocks: func(clientRepo *mock.MockClientRepository, cache *mock.MockCacheRepository) {
				clientRepo.EXPECT().CreateClient(gomock.Any(), gomock.Eq(clientInput)).Return(nil, domain.ErrConflictingData)
			},
			input: registerTestedInput{
				clientInput,
			},
			expected: registerExpectedOutput{
				client: nil,
				err:    domain.ErrConflictingData,
			},
		},
		{
			desc: "Fail_InternalError",
			mocks: func(clientRepo *mock.MockClientRepository, cache *mock.MockCacheRepository) {
				clientRepo.EXPECT().CreateClient(gomock.Any(), gomock.Eq(clientInput)).Return(nil, domain.ErrInternal)
			},
			input: registerTestedInput{
				clientInput,
			},
			expected: registerExpectedOutput{
				client: nil,
				err:    domain.ErrInternal,
			},
		},
		{
			desc: "Fail_SetCache",
			mocks: func(clientRepo *mock.MockClientRepository, cache *mock.MockCacheRepository) {
				clientRepo.EXPECT().CreateClient(gomock.Any(), gomock.Eq(clientInput)).Return(clientOutput, nil)
				cache.EXPECT().Set(gomock.Any(), gomock.Eq(cacheKey), gomock.Eq(userSerialized), gomock.Eq(ttl)).Return(domain.ErrInternal)
			},
			input: registerTestedInput{
				clientInput,
			},
			expected: registerExpectedOutput{
				client: nil,
				err:    domain.ErrInternal,
			},
		},
		{
			desc: "Fail_DeleteCacheByPrefix",
			mocks: func(clientRepo *mock.MockClientRepository, cache *mock.MockCacheRepository) {
				clientRepo.EXPECT().CreateClient(gomock.Any(), gomock.Eq(clientInput)).Return(clientOutput, nil)
				cache.EXPECT().Set(gomock.Any(), gomock.Eq(cacheKey), gomock.Eq(userSerialized), gomock.Eq(ttl)).Return(nil)
				cache.EXPECT().DeleteByPrefix(gomock.Any(), gomock.Eq("clients:*")).Return(domain.ErrInternal)
			},
			input: registerTestedInput{
				clientInput,
			},
			expected: registerExpectedOutput{
				client: nil,
				err:    domain.ErrInternal,
			},
		},
	}

	for _, tc := range testCases {
		// tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			// // TODO: fix race condition to enable parallel testing
			// t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			clientRepo := mock.NewMockClientRepository(ctrl)
			cache := mock.NewMockCacheRepository(ctrl)

			tc.mocks(clientRepo, cache)

			clientService := service.NewClientService(clientRepo, cache)

			client, err := clientService.Register(ctx, tc.input.client)
			assert.Equal(t, tc.expected.err, err, "Error mismatch")
			assert.Equal(t, tc.expected.client, client, "Client mismatch")
		})
	}
}
