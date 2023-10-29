set(target_name geometry-c)

add_library(${target_name} SHARED
  libgeometry-c/geometry.h
  libgeometry-c/Circle.cpp
  libgeometry-c/Circle.h
  libgeometry-c/Point.h
  libgeometry-c/Shape.cpp
  libgeometry-c/Shape.h
)

include(CompileOptions)
set_compile_options(${target_name})

target_include_directories(
  ${target_name}
  PUBLIC
  ${CMAKE_CURRENT_SOURCE_DIR}
)

target_link_libraries(${target_name} PUBLIC geometry)

add_custom_command(
  OUTPUT ${CMAKE_SOURCE_DIR}/client/geometry.out
  COMMAND make build include_path=${CMAKE_CURRENT_LIST_DIR} slib_path=$<TARGET_FILE_DIR:${target_name}> libname=${target_name}
  DEPENDS ${CMAKE_SOURCE_DIR}/client/cmd/main.go ${CMAKE_SOURCE_DIR}/client/geometry/geometry.go
  WORKING_DIRECTORY ${CMAKE_SOURCE_DIR}/client
)

add_custom_target(
  build_geometry
  ALL
  DEPENDS ${target_name} ${CMAKE_SOURCE_DIR}/client/geometry.out
)
