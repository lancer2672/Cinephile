package converter

import (
	"github.com/lancer2672/cinephile/main/internal/domain/model"
	"github.com/lancer2672/cinephile/main/internal/interface/grpc/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToProtoMovie(m *model.Movie) *pb.Movie {
	if m == nil {
		return nil
	}

	return &pb.Movie{
		Id:          m.ID,
		Title:       m.Title,
		Duration:    m.Duration, // giữ đơn vị giây
		AgeLimit:    m.AgeLimit,
		Director:    m.Director,
		Actors:      m.Actors,
		ReleaseDate: timestamppb.New(m.ReleaseDate),
		Description: m.Description,
		Trailer:     m.Trailer,
		Country:     m.Country,
		Genre:       m.Genre,
		Poster:      m.Poster,
		Status:      m.Status,
		CreatedAt:   timestamppb.New(m.CreatedAt),
		UpdatedAt:   timestamppb.New(m.UpdatedAt),
	}
}
