require './lib/memo_definition'
require './lib/memo_node'

class MemoManager
  def initialize(memo)
    @memoObj = parse(memo)
  end

  def get_node(line, parent)
    case line
    when /^###/
      title, attribute = MemoDefinition.get_from_level_3(line)
      MemoNode.new(title, line, parent, attribute, 3)
    when /^##/
      title, attribute = MemoDefinition.get_from_level_2(line)
      MemoNode.new(title, line, parent, attribute, 2)
    when /^#/
      title, attribute = MemoDefinition.get_from_level_1(line)
      MemoNode.new(title, line, parent, attribute, 1)
    end
  end

  def generate_tree(each_line, parent, index)
    if each_line[index] =~ /^#/
      child = get_node(each_line[index], parent)

      # レベルが上がると親が変わるので親まで遡る
      if parent.level >= child.level
        while parent.level >= child.level do parent = parent.parent end
      end
      parent.append_node(child)
    else
      parent.append_text(each_line[index])
      child = parent
    end

    if index < each_line.size()-1
      generate_tree(each_line, child, index+1)
    end
  end

  def parse(memo)
    each_line = memo.split(/\R/)
    if each_line[0] =~ /^#\s+/
      parent = get_node(each_line[0], nil)
      generate_tree(each_line, parent, 1)
    else
      raise 'Missing Format: First must memo title'
    end

    return parent
  end

  def get_level_1()
    return [@memoObj]
  end

  def get_level_2()
    return @memoObj.children
  end

  def get_level_3()
    return @memoObj.children.map do |child|
      child.children
    end.flatten()
  end

  def self.decode(parent = nil)
    if parent == nil then parent = @memoObj end

    text = parent.text
    for child in parent.children
      text += "\n"+decode(child)
    end

    return text
  end
end
