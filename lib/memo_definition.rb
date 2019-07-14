class MemoDefinition
  def self.level_2_options(opts)
    case opts
    when 'yet'
      ['WIP', 'REVIEW', 'WAIT']
    end
  end

  def self.get_from_level_3(line)
    case line
    when /^###\s\[WORK\]/
      [line.match(/(###\s\[WORK\]\s+(.+))/)[2], 'WORK']
    when /^###\s\[REFS\]/
      ['', 'REFS']
    when /^###\s\[WHAT\]/
      ['', 'WHAT']
    when /^###\s\[WHO\]/
      ['', 'WHO']
    when /^###\s\[WHEN\]/
      ['', 'WHEN']
    when /^###\s\[WHERE\]/
      ['', 'WHERE']
    when /^###\s\[WHY\]/
      ['', 'WHY']
    when /^###\s\[HOW\]/
      ['', 'HOW']
    else
      raise 'missing level 3 attribute'
    end
  end

  def self.get_from_level_2(line)
    case line
    when /^##\s\[REVIEW\]/
      [line.match(/(##\s\[REVIEW\]\s+(.+))/)[2], 'REVIEW']
    when /^##\s\[WIP\]/
      [line.match(/(##\s\[WIP\]\s+(.+))/)[2], 'WIP']
    when /^##\s\[WAIT\]/
      [line.match(/(##\s\[WAIT\]\s+(.+))/)[2], 'WAIT']
    when /^##\s\[DONE\]/
      [line.match(/(##\s\[DONE\]\s+(.+))/)[2], 'DONE']
    else
      raise 'missing level 2 attribute'
    end
  end

  def self.get_from_level_1(line)
    case line
    when /^#\s\[TOP\]/
      [line.match(/(#\s\[TOP\]\s+(.+))/)[2], 'TOP']
    else
      raise 'missing level 1 attribute'
    end
  end
end
