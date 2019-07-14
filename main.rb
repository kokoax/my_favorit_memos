require './lib/memo_manager'

def get_title(mm)
  mm.get_level_2().map do |item|
    item.get_title()
  end
end

def get_title_with_index(mm)
  mm.get_level_2().map.with_index do |item, index|
    "#{index}: #{item.get_title()}"
  end
end

def get_node_from_index(mm, index)
  MemoManager.decode(mm.get_level_2()[index])
end

def get_node_from_title(mm, title)
  MemoManager.decode(
    (mm.get_level_2().select do |item|
    item.get_title() == title
  end)[0])
end

def edit_node(mm, node)
end

def main(filepath, opts)
  memo = File.open(filepath, "r").read
  mm = MemoManager.new(memo)

  # puts get_title_with_index(mm)
  # puts get_node(mm, '[WIP] CI直す前にリリースしちゃった問題')
  # puts get_node_from_index(mm, 1)
end

main('./memos.md', ARGV)
