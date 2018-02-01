Rails.application.routes.draw do
  resources :users
  resources :lectures, only: [:index, :new, :create, :destroy]
  resources :researchs, only: [:index, :new, :create, :destroy]

  get '/login', to: 'sessions#login_form'
  post '/login', to: 'sessions#login'
  post '/logout', to: 'sessions#logout'
  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
end
