Rails.application.routes.draw do
  root 'static_pages#index'

  resources :users do
    member do
      get :edit_pass
      post :update_pass
    end
  end
  resources :lectures do
    member do
      get :download
    end
  end
  resources :researches, only: [:index, :new, :create, :destroy] do
    member do
      get :download
    end
  end
  resources :equipments

  get '/login', to: 'sessions#login_form'
  post '/login', to: 'sessions#login'
  delete '/logout', to: 'sessions#logout'

  get '/activity', to: 'static_pages#activity'
  get '/publication', to: 'static_pages#publication'
  get '/job', to: 'static_pages#job'
  get '/link', to: 'static_pages#link'
  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
end
