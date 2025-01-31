package vehiclepositions

import (
	"testing"
	"time"

	gtfsRt_vehiclepositions "github.com/CanalTP/forseti/internal/gtfsRt_vehiclepositions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_NewVehiclePosition(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)
	location, err := time.LoadLocation("Europe/Paris")
	require.Nil(err)
	date, err := time.ParseInLocation("2006-01-02", "2021-02-22", location)
	require.Nil(err)

	vp, err := NewVehiclePosition(651969, "vehicle_journey:STS:651969-1", date, 45.398613, -71.90111, 0, 0)
	assert.Nil(err)
	require.NotNil(vp)
}

func Test_UpdateVehiclePosition(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	location, err := time.LoadLocation("Europe/Paris")
	require.Nil(err)

	vehiclePositions := VehiclePositions{}

	vGtfsRt := gtfsRt_vehiclepositions.VehicleGtfsRt{VehicleID: "52103", StopId: "1280", Label: "52103", Time: 1621900800,
		Speed: 0, Bearing: 0, Route: "1", Trip: "651970", Latitude: 45.9999, Longitude: -71.90111, Occupancy: 0}
	changeGtfsRt := gtfsRt_vehiclepositions.VehicleGtfsRt{VehicleID: "52103", StopId: "1280", Label: "52103",
		Time: 1621900800, Speed: 11, Bearing: 254, Route: "1", Trip: "651970", Latitude: 46.0000,
		Longitude: -72.0000, Occupancy: 0}

	// Update vehiclePosition with no map cehiclelocations
	vehiclePositions.UpdateVehiclePosition(1, changeGtfsRt, location)
	require.Nil(vehiclePositions.vehiclePositions)

	// Create vehiclePosition from existing data
	vp := createVehiclePositionFromDataSource(1, vGtfsRt, location)
	//t.Log("DATE: ", vp.DateTime)
	vehiclePositions.AddVehiclePosition(vp)
	require.NotNil(vehiclePositions.vehiclePositions)
	assert.Equal(len(vehiclePositions.vehiclePositions), 1)

	// Update vehiclePosition with existing data
	vehiclePositions.UpdateVehiclePosition(1, changeGtfsRt, location)
	assert.Equal(vehiclePositions.vehiclePositions[1].Latitude, float32(46.0000))
	assert.Equal(vehiclePositions.vehiclePositions[1].Longitude, float32(-72.0000))
	assert.Equal(vehiclePositions.vehiclePositions[1].Bearing, float32(254))
	assert.Equal(vehiclePositions.vehiclePositions[1].Speed, float32(11))
}
