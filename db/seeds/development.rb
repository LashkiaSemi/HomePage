ActiveRecord::Base.logger = Logger.new(STDOUT)
ActiveRecord::Base.transaction do
    ## seedで入れたいデータを挿入する処理を記述
    Importer::NewsImporter.import
end