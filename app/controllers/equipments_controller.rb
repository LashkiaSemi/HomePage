class EquipmentsController < ApplicationController
  before_action :is_user_admin, except: [:index]

  def index
    @equipment = Equipment.all.order(:tag) 
  end

  def new
    @eqipment = Equipment.new
    @equipment.build_tag
  end
  
  def create
    @equipment = Equipment.new(equipment_params)
    if @equipment.save
      flash[:success] = "備品を登録しました。"
      redirect_to equipments_url
    else
      flash[:danger] = "備品の登録に失敗しました。"
      render action: :new
    end
  end
  
  def edit
    @equipment = Equipment.find(params[:id])
  end

  def update
    @equipment = Equipment.find(params[:id])
    if @equipment.update(equipment_params)
      flash[:success] = "備品を編集しました。"
      redirect_to equipment_url(@equipment)
    else
      flash[:danger] = "備品の編集に失敗しました。"
      render action: :edit
    end
  end
  
  def destroy
    Equipment.find(params[:id]).destroy
  end  
end
