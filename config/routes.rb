Rails.application.routes.draw do
  root 'static_pages#index'

  resources :users do
    member do
      get :edit_pass
      post :update_pass
    end
  end
  resources :lectures, only: [:index, :new, :create, :destroy]
  resources :researchs, only: [:index, :new, :create, :destroy]

  get '/login', to: 'sessions#login_form'
  post '/login', to: 'sessions#login'
  post '/logout', to: 'sessions#logout'

  get '/activitiy', to: 'static_pages#activitiy'
  get '/equipment', to: 'static_pages#equipment'
  get '/publication', to: 'static_pages#publication'
  get '/job', to: 'static_pages#job'
  get '/link', to: 'static_pages#link'
  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
end
