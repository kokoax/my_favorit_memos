class MemoNode
  attr_reader :title
  attr_reader :text
  attr_reader :parent
  attr_reader :children
  attr_reader :attr
  attr_reader :level

  def initialize(title, text, parent, attr, level)
    @title = title
    @text = text
    @parent = parent
    @children = []
    @attr = attr
    @level = level
  end

  def get_title()
    "[#{@attr}] #{@title}"
  end

  def append_node(node)
    @children.push(node)
  end

  def append_text(text)
    @text << "\n" << text
  end
end
