class EquipmentsController < ApplicationController
  before_action :is_user_admin, except: [:index]

  def index
    @equipments = Equipment.all
  end

  def new
    if authority_admin?
      @equipment = Equipment.new
    else
      flash[:warning] = "権限がありません。"
      redirect_to equipments_url
    end
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
    if authority_admin?
      @equipment = Equipment.find(params[:id])
      @tags = Tag.all
    else
      flash[:warning] = "権限がありません。"
      redirect_to equipments_url
    end
  end

  def update
    @equipment = Equipment.find(params[:id])
    if @equipment.update(equipment_params)
      flash[:success] = "備品を編集しました。"
      redirect_to equipments_url(@equipment)
    else
      flash[:danger] = "備品の編集に失敗しました。"
      render action: :edit
    end
  end
  
  def destroy
    Equipment.find(params[:id]).destroy
    redirect_to equipments_url
  end

  private
    def equipment_params
      params.require(:equipment).permit(:name, :num, :note, :tag_id)
    end
end
