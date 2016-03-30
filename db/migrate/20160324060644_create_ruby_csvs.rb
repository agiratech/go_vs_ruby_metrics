class CreateRubyCsvs < ActiveRecord::Migration
  def change
    create_table :ruby_csvs do |t|
      t.string :policy_id
      t.string :state_code
      t.string :county
      t.string :eq_site_limit
      t.string :hu_site_limit
      t.string :fl_site_limit
      t.string :fr_site_limit
      t.string :tiv_2011
      t.string :tiv_2012
      t.string :eq_site_deductible
      t.string :hu_site_deductible
      t.string :fl_site_deductible
      t.string :fr_site_deductible
      t.string :point_latitude
      t.string :point_longitude
      t.string :line
      t.string :construction
      t.string :point_granularity

      t.timestamps null: false
    end
  end
end
