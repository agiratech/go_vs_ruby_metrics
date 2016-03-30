# encoding: UTF-8
# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# Note that this schema.rb definition is the authoritative source for your
# database schema. If you need to create the application database on another
# system, you should be using db:schema:load, not running all the migrations
# from scratch. The latter is a flawed and unsustainable approach (the more migrations
# you'll amass, the slower it'll run and the greater likelihood for issues).
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema.define(version: 20160324060644) do

  # These are extensions that must be enabled in order to support this database
  enable_extension "plpgsql"

  create_table "ruby_csvs", force: :cascade do |t|
    t.string   "policy_id"
    t.string   "state_code"
    t.string   "county"
    t.string   "eq_site_limit"
    t.string   "hu_site_limit"
    t.string   "fl_site_limit"
    t.string   "fr_site_limit"
    t.string   "tiv_2011"
    t.string   "tiv_2012"
    t.string   "eq_site_deductible"
    t.string   "hu_site_deductible"
    t.string   "fl_site_deductible"
    t.string   "fr_site_deductible"
    t.string   "point_latitude"
    t.string   "point_longitude"
    t.string   "line"
    t.string   "construction"
    t.string   "point_granularity"
    t.datetime "created_at",         null: false
    t.datetime "updated_at",         null: false
  end

end
