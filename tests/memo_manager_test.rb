require 'test/unit'

require './lib/memo_manager'

class MemoManagerTest < Test::Unit::TestCase
  def test_parse
    data = '# [TOP] sample'
    memo = MemoManager.new(data)

    expect = ''

    res = memo.parse(data)
    assert_equal(res, expect)
  end
end
