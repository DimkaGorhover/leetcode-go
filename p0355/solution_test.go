package p0355

import (
	"reflect"
	"testing"
)

type twitterAssertions struct {
	t       *testing.T
	twitter *Twitter
}

func (t twitterAssertions) assertNewsFeed(userId int, expected []int) {
	actual := t.twitter.GetNewsFeed(userId)
	if !reflect.DeepEqual(expected, actual) {
		t.t.Errorf(`twitter.GetNewsFeed(%d); expected=%v; actual=%v`, userId, expected, actual)
	}
}

func TestConstructor(t *testing.T) {
	t.Parallel()

	t.Run(`test_001`, func(t *testing.T) {
		twitter := Constructor()
		assertions := twitterAssertions{t, &twitter}

		twitter.PostTweet(1, 5)
		twitter.PostTweet(1, 3)

		assertions.assertNewsFeed(1, []int{3, 5})
	})

	t.Run(`test_002`, func(t *testing.T) {
		twitter := Constructor()
		assertions := twitterAssertions{t, &twitter}

		twitter.Follow(1, 2)
		twitter.Follow(1, 3)
		twitter.Follow(1, 4)
		twitter.PostTweet(4, 4)
		twitter.PostTweet(3, 3)
		twitter.PostTweet(2, 2)
		twitter.PostTweet(1, 1)

		assertions.assertNewsFeed(1, []int{4, 3, 2, 1})
		assertions.assertNewsFeed(2, []int{2})
	})

	t.Run(`test_003`, func(t *testing.T) {
		twitter := Constructor()
		assertions := twitterAssertions{t, &twitter}

		twitter.PostTweet(1, 6765)
		twitter.PostTweet(5, 671)
		twitter.PostTweet(3, 2868)
		twitter.PostTweet(4, 8148)
		twitter.PostTweet(4, 386)
		twitter.PostTweet(3, 6673)
		twitter.PostTweet(3, 7946)
		twitter.PostTweet(3, 1445)
		twitter.PostTweet(4, 4822)
		twitter.PostTweet(1, 3781)
		twitter.PostTweet(4, 9038)
		twitter.PostTweet(1, 9643)
		twitter.PostTweet(3, 5917)
		twitter.PostTweet(2, 8847)
		twitter.Follow(1, 3)
		twitter.Follow(1, 4)
		twitter.Follow(4, 2)
		twitter.Follow(4, 1)
		twitter.Follow(3, 2)
		twitter.Follow(3, 5)
		twitter.Follow(3, 1)
		twitter.Follow(2, 3)
		twitter.Follow(2, 1)
		twitter.Follow(2, 5)
		twitter.Follow(5, 1)
		twitter.Follow(5, 2)
		assertions.assertNewsFeed(1, []int{5917, 9643, 9038, 3781, 4822, 1445, 7946, 6673, 386, 8148})
		assertions.assertNewsFeed(2, []int{8847, 5917, 9643, 3781, 1445, 7946, 6673, 2868, 671, 6765})
		assertions.assertNewsFeed(3, []int{8847, 5917, 9643, 3781, 1445, 7946, 6673, 2868, 671, 6765})
		assertions.assertNewsFeed(4, []int{8847, 9643, 9038, 3781, 4822, 386, 8148, 6765})
		assertions.assertNewsFeed(5, []int{8847, 9643, 3781, 671, 6765})
	})
}
