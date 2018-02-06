class Document < ApplicationRecord
    validates :name, presence: true
    validates :path, presence: true
end
