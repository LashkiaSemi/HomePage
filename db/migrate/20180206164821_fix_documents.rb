class FixDocuments < ActiveRecord::Migration[5.1]
  def change
    rename_column :documents, :boolean, :activation
    add_column :documents, :comments, :string
  end
end
