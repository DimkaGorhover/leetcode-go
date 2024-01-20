package lc0919

import (
	. "github.com/DimkaGorhover/leetcode-go/pkg/common"
	"reflect"
	"testing"
)

func TestCBTInserter_Get_root(t *testing.T) {
	t.Parallel()

	t.Run(`test_001`, func(t *testing.T) {

		assertEqualsInt := func(expected, actual int) {
			if !reflect.DeepEqual(expected, actual) {
				t.Errorf(`expected: %v; Actual = %v`, expected, actual)
			}
		}

		assertEqualsTreeNode := func(expected, actual *TreeNode) {
			if !expected.Equals(actual) {
				t.Errorf(`expected: %v; Actual = %v`, expected, actual)
			}
		}

		cBTInserter := Constructor(NewTreeNodeBuilder().AddAll(1, 2, 3, 4, 5, 6).Build())
		assertEqualsInt(3, cBTInserter.Insert(7))
		assertEqualsInt(4, cBTInserter.Insert(8))
		assertEqualsTreeNode(NewTreeNodeBuilder().AddAll(1, 2, 3, 4, 5, 6, 7, 8).Build(), cBTInserter.Get_root())

		//val cBTInserter = new CBTInserter(TreeNode.of(1, 2, 3, 4, 5, 6));
		//assertThat(cBTInserter.insert(7)).isEqualTo(3);
		//assertThat(cBTInserter.insert(8)).isEqualTo(4);
		//assertThat(cBTInserter.get_root()).isEqualTo(TreeNode.of(1, 2, 3, 4, 5, 6, 7, 8));
	})

	t.Run(`test_002`, func(t *testing.T) {
		//val cBTInserter = new CBTInserter(TreeNode.of(1, 2));
		//assertThat(cBTInserter.insert(3)).isEqualTo(1);
		//assertThat(cBTInserter.insert(4)).isEqualTo(2);
		//assertThat(cBTInserter.get_root()).isEqualTo(TreeNode.of(1, 2, 3, 4));
	})
}
