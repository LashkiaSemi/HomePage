class News < ApplicationRecord
    validates :title, presence: true
    validates :state, presence: true
    validates :time, presence: true

    enum state: {
        top: 0,
        published: 1,
        privated: 2
    }
end
