require 'test/unit'

require './lib/memo_node'

class MemoNodeTest < Test::Unit::TestCase
  def test_initialize()
    res = MemoNode.new('sample', [], 'attr')
    expect = MemoNode.new('sample', [], 'attr')

    assert_equal(res, expect)
  end
end


