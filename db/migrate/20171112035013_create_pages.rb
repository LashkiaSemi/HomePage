class CreatePages < ActiveRecord::Migration[5.1]
  def change
    create_table :pages do |t|
      t.string  :title
      t.string  :contents
      t.timestamps
    end
  end
end
