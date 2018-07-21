require 'csv'

module Importer
  module NewsImporter
    CSV_PATH = Rails.root.join('db', 'seeds', 'csv', 'older_news.csv')

    class << self
      def import(csv_path = CSV_PATH)
        CSV.foreach(csv_path, headers: true) do |row|
          news = News.find_or_initialize_by(id: row['id'])
          news.attributes = build_news_attributes(row)
          news.save if news.changed?
        end
      end

      private

      def build_news_attributes(row)
        {
          id: row['id'],
          title: row['title'],
          state: row['state'],
          time: row['time']
        }
      end
    end
  end
end