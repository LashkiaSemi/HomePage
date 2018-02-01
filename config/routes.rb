Rails.application.routes.draw do
  root 'static_pages#index'

  
  get '/activitiy', to: 'static_pages#activitiy'
  get '/equipment', to: 'static_pages#equipment'
  get '/publication', to: 'static_pages#publication'
  get '/job', to: 'static_pages#job'
  get '/link', to: 'static_pages#link'

  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
end
