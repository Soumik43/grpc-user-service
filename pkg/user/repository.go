package user

import (
	"fmt"
	"sync"

	pb "github.com/Soumik43/grpc-user-service/api/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Repository interface {
	GetUser(id int32) (*pb.User, error)
	GetUsers(ids []int32) (*pb.Users, error)
}

type InMemoryUserRepository struct {
	// mutex protects the users map for concurrent access when writing to avoid race conditions (future usecase)
	mu    sync.Mutex
	Users map[int32]*pb.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	repo := &InMemoryUserRepository{
		Users: make(map[int32]*pb.User),
	}

	// Pre-populate with some users
	repo.Users[1] = &pb.User{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true}
	repo.Users[2] = &pb.User{Id: 2, Fname: "Shimmy", City: "NYC", Phone: 9876543210, Height: 5.2, Married: false}
	repo.Users[3] = &pb.User{Id: 3, Fname: "Jane", City: "HongKong", Phone: 1234509876, Height: 5.1, Married: true}
	repo.Users[4] = &pb.User{Id: 4, Fname: "Soumik", City: "Mumbai", Phone: 1234509876, Height: 6.0, Married: true}
	repo.Users[5] = &pb.User{Id: 5, Fname: "Suraj", City: "Mumbai", Phone: 1234509876, Height: 5.6, Married: false}
	repo.Users[6] = &pb.User{Id: 6, Fname: "Dalvi", City: "HKL", Phone: 1234509876, Height: 5.3, Married: true}
	repo.Users[7] = &pb.User{Id: 7, Fname: "Stooper", City: "CCU", Phone: 1234509876, Height: 5.2, Married: false}
	repo.Users[8] = &pb.User{Id: 8, Fname: "Mack", City: "BLK", Phone: 1234509876, Height: 5.9, Married: false}
	repo.Users[9] = &pb.User{Id: 9, Fname: "Shahin", City: "Canada", Phone: 1234509876, Height: 5.16, Married: true}
	repo.Users[10] = &pb.User{Id: 10, Fname: "Laura", City: "Austin", Phone: 4928394032, Height: 5.2, Married: false}

	return repo
}

func (r *InMemoryUserRepository) GetUser(id int32) (*pb.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check negative UserId
	if id < 0 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Invalid User ID: %d", id))
	}

	user, exists := r.Users[id]
	if !exists {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("User with id %d does not exist", id))
	}

	return user, nil
}

func (r *InMemoryUserRepository) GetUsers(ids []int32) (*pb.Users, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check empty UserIdList
	if len(ids) == 0 {
		return nil, status.Error(codes.InvalidArgument, "No User IDs provided")
	}

	userList := &pb.Users{}
	for _, id := range ids {
		if user, exists := r.Users[id]; exists {
			userList.Users = append(userList.Users, user)
		}
	}

	return userList, nil
}
