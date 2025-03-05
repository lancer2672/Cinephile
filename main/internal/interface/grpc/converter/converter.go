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
		MovieId:            int64(m.ID),
		Title:              m.Name,
		Tags:               m.Tags,
		Duration:           int64(m.Duration.Seconds()), // Chuyển `time.Duration` sang giây
		AgeLimit:           int32(m.AgeLimit),
		Director:           m.Director,
		Actors:             m.Actors,
		ReleaseDate:        timestamppb.New(m.ReleaseDate), // Chuyển `time.Time` sang `Timestamp`
		ContentDescription: m.ContentDescription,
		Trailer:            m.Trailer,
		Poster:             m.Poster,
	}
}
