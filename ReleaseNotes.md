v0.3.2
- ...
- removed ConsoleProducer clone used during emitter refactoring
- improved error messages if template name doesn't exist
- 

v0.3.1
- Added $JR_HOME 
- Added $PATH and $JR_HOME to jrconfig search
- Made templates subdir of $JR_HOME and got rid of --templateDir option
- merged jrconfig.json conf and cli conf in jr template run
- Minor refactorings to interface Producer
- Added UK localization
- Added ES localization
- Added back jr commands as jr template run|list|show
- Added stock symbol function
- Fixed user avro schema
- Started http server refactoring -> uncompleted

v0.3.0
- Added emitters, producer, functions and templates as resources
- Added emitter concept and refactored run/template run to use an emitter
- Added E2E example with relations (shoes)
- Added server with REST apis
- Added preload
- Added multithread support for emitters
- Added S3 producer
- Added Elastic Producer
- Added Mongo producer
- Added relations between templates (random_n_v_from_list, add_v_to_list random_v_from_list)
- Fixed seed in uuid
- Added French localization
- Switched to ubi images in docker
- Added avro files for all old Datagen templates
- Added function counter to function list/man
- Added producer list
- Added jrconfig to configure everything (emitters, global) in a single file
- Default key template for Kafka changed to null (no key)

v0.1.9
- improved generate.go to generate *.avsc
- updated release notes
- added date_between, dates_between and birthdate functions
- added past, future, recent and soon functions. Fixes #53
- added imei phone function
- added land_prefix italian zip and phone country code
- added city and city at. Added real regex zip codes for Italian cities
- updated zip codes file
- added country_code, country_code_at, land_prefix, land_prefix_at, country_at functions
- added mobile_phone, mobile_phone_at functions. Added several regexs for many countries

v0.1.8

- Added ISO 3166 country function
- Added credit card generator, fixed other stuff
- Added cusip code
- Added ssn and cf skeleton
- Added italian codice fiscale
- if locale doesn't exist, use US
- Added account and amount finance functions
- Added output in man description of some functions
- Added swift function
- Added bitcoin function
- Added cardCVV function
- Added ethereum address
- Added image and image_of functions. Fix #52

v0.1.7
- added go:generator for automatic registry generation
- added mac workflow to github
- fixed template names

v0.1.6
- internal refactoring
- fixed gogen-avro path
