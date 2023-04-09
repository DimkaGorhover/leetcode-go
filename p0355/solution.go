package p0355

import (
	"container/heap"
	"fmt"
)

// ============================================================================
// Twitter

type Twitter struct {
	userFactory *UserFactory
	users       map[int]*User
}

func Constructor() Twitter {
	return Twitter{
		userFactory: NewUserFactory(),
		users:       make(map[int]*User),
	}
}

func (t *Twitter) getOrCreateUser(id int) *User {
	user, found := t.users[id]
	if !found {
		user = t.userFactory.NewUser(id)
		t.users[id] = user
	}
	return user
}

func (t *Twitter) PostTweet(userId int, tweetId int) {
	t.getOrCreateUser(userId).AddTweet(tweetId)
}

func (t *Twitter) GetNewsFeed(userId int) []int {
	return t.getOrCreateUser(userId).GetNewsFeed()
}

func (t *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	follower := t.getOrCreateUser(followerId)
	followee := t.getOrCreateUser(followeeId)
	follower.Follow(followee)
}

func (t *Twitter) Unfollow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	follower := t.getOrCreateUser(followerId)
	follower.Unfollow(followeeId)
}

// ============================================================================
// UserFactory

type UserFactory struct {
	tweetFactory *TweetFactory
}

func NewUserFactory() *UserFactory {
	return &UserFactory{
		tweetFactory: NewTweetFactory(),
	}
}

func (f *UserFactory) NewUser(id int) *User {
	return &User{
		tweetFactory: f.tweetFactory,
		id:           id,
		following:    make(map[int]*User),
	}
}

// ============================================================================
// User

type User struct {
	tweetFactory  *TweetFactory
	id            int
	following     map[int]*User
	firstFollowee *User
	headTweet     *Tweet
	tweetsCount   int
}

func (u *User) String() string {
	return fmt.Sprintf(`User{id=%d, tweetsCount=%d}`, u.id, u.tweetsCount)
}

func (u *User) Follow(followee *User) {
	if followee == nil || u.id == followee.id {
		return
	}
	u.following[followee.id] = followee
	if u.firstFollowee == nil {
		u.firstFollowee = followee
	}
}

func (u *User) Unfollow(id int) {
	delete(u.following, id)
}

func (u *User) AddTweet(tweetId int) {
	u.headTweet = u.tweetFactory.NewTweet(tweetId, u.headTweet)
	u.tweetsCount++
}

func (u *User) GetNewsFeed() []int {
	limit := 10

	if len(u.following) == 0 {
		limit = minInt(limit, u.tweetsCount)
		arr := make([]int, limit)
		t := u.headTweet
		for i := 0; t != nil && limit > 0; i++ {
			arr[i] = t.id
			t = t.next
			limit--
		}

		return arr
	}

	if len(u.following) == 1 {
		limit = minInt(limit, u.tweetsCount+u.firstFollowee.tweetsCount)
		arr := make([]int, limit)
		i := 0
		t1 := u.headTweet
		t2 := u.firstFollowee.headTweet

		for t1 != nil && t2 != nil && limit > 0 {
			if t1.time > t2.time {
				arr[i] = t1.id
				t1 = t1.next
			} else {
				arr[i] = t2.id
				t2 = t2.next
			}
			i++
			limit--
		}

		for t1 != nil && limit > 0 {
			arr[i] = t1.id
			t1 = t1.next
			i++
			limit--
		}

		for t2 != nil && limit > 0 {
			arr[i] = t2.id
			t2 = t2.next
			i++
			limit--
		}

		return arr
	}

	tweetsHeap := NewTweetsHeap()
	tweetsHeap.Push(u.headTweet)
	capacity := u.tweetsCount
	for _, followee := range u.following {
		tweetsHeap.Push(followee.headTweet)
		capacity += followee.tweetsCount
	}

	limit = minInt(limit, capacity)
	arr := make([]int, limit)
	i := 0
	for limit > 0 && !tweetsHeap.IsEmpty() {
		tweet := tweetsHeap.Pop()
		arr[i] = tweet.id
		tweetsHeap.Push(tweet.next)
		i++
		limit--
	}

	return arr
}

// ============================================================================
// TweetFactory

type TweetFactory struct {
	tweetTimer uint64
}

func NewTweetFactory() *TweetFactory {
	return &TweetFactory{}
}

func (f *TweetFactory) NewTweet(id int, next *Tweet) *Tweet {
	time := f.tweetTimer
	f.tweetTimer++
	return &Tweet{
		id:   id,
		next: next,
		time: time,
	}
}

// ============================================================================
// Tweet

type Tweet struct {
	id   int
	next *Tweet
	time uint64
}

// ============================================================================
// TweetsHeap

type TweetsHeap struct {
	t *TweetsHeapContainer
}

func NewTweetsHeap() *TweetsHeap {
	c := &TweetsHeapContainer{
		arr:    make([]*Tweet, 0),
		length: 0,
	}
	return &TweetsHeap{t: c}
}

func (t *TweetsHeap) Push(x *Tweet) {
	if x != nil {
		heap.Push(t.t, x)
	}
}
func (t *TweetsHeap) Pop() *Tweet {
	return heap.Pop(t.t).(*Tweet)
}
func (t *TweetsHeap) IsEmpty() bool {
	return t.t.Len() == 0
}

// ============================================================================
// TweetsHeapContainer

type TweetsHeapContainer struct {
	arr    []*Tweet
	length int
}

func (t *TweetsHeapContainer) Len() int {
	return t.length
}

func (t *TweetsHeapContainer) Less(i, j int) bool {
	return t.arr[i].time > t.arr[j].time
}

func (t *TweetsHeapContainer) Swap(i, j int) {
	t.arr[i], t.arr[j] = t.arr[j], t.arr[i]
}

func (t *TweetsHeapContainer) Push(x interface{}) {
	tweet := x.(*Tweet)
	if len(t.arr) > t.length {
		t.arr[t.length] = tweet
	} else {
		t.arr = append(t.arr, tweet)
	}
	t.length++
}

func (t *TweetsHeapContainer) Pop() interface{} {
	if t.length == 0 {
		return nil
	}
	t.length--
	return t.arr[t.length]
}

// ============================================================================
// Utils

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
