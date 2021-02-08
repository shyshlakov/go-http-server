package service

// import (
// 	"fmt"
// 	"testing"
// 	"time"

// 	"github.com/shyshlakov/go-http-server/persistence/model"
// 	"github.com/shyshlakov/go-http-server/persistence/repo"

// 	"github.com/golang-migrate/migrate/v4"
// 	"github.com/golang-migrate/migrate/v4/database/postgres"
// 	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
// 	"github.com/jmoiron/sqlx"
// 	"github.com/ory/dockertest/v3"
// 	"github.com/rs/zerolog/log"
// 	gouuid "github.com/satori/go.uuid"
// 	"github.com/shyshlakov/go-http-server/config"
// 	repoPostgres "github.com/shyshlakov/go-http-server/persistence/repo/postgres"

// 	// "github.com/shyshlakov/go-http-server/services/migration/schema"
// 	// "github.com/shyshlakov/go-http-server/services/persistence/consts"
// 	// "github.com/shyshlakov/go-http-server/services/persistence/repo"
// 	// "github.com/shyshlakov/go-http-server/services/restapi"
// 	// "github.com/shyshlakov/go-http-server/services/state"
// 	"github.com/stretchr/testify/mock"
// 	"github.com/stretchr/testify/suite"
// )

// type KontomatikTestSuite struct {
// 	suite.Suite
// 	db       *sqlx.DB
// 	pool     *dockertest.Pool
// 	resource *dockertest.Resource
// 	srv      *AppService
// 	commands map[string]*CommandMockObject
// 	cfg      *config.Config
// 	repo     repo.Repo
// }

// type CommandMockObject struct {
// 	mock.Mock
// }

// func (c *CommandMockObject) StartImportData(params map[string]string) (string, error) {
// 	args := c.Called(params)
// 	return args.String(0), args.Error(1)
// }

// func (c *CommandMockObject) GetDataByCommandID(reqID gouuid.UUID, sessionID, signature, commID string) (*kontomatik.ReplyData, error) {
// 	args := c.Called(reqID, sessionID, signature, commID)
// 	return args.Get(0).(*kontomatik.ReplyData), args.Error(1)
// }

// func (c *CommandMockObject) ProcessData(reqID gouuid.UUID, data *kontomatik.ReplyData, resp *model.TargetReportData) error {
// 	args := c.Called(reqID, data)
// 	return args.Error(0)
// }

// const testOverallData = `
// <target name="MBank" officialName="mBank" institution="MBank">
// 	<owners>
// 		<owner>
// 			<name>Jan Kowalski</name>
// 			<address>Kwiatowa 8/97 Warszawa  PL</address>
// 			<polishPesel>84011806651</polishPesel>
// 			<phone>+48612***078</phone>
// 			<email>jankowalski@xgmail.com</email>
// 			<citizenship>PL</citizenship>
// 			<personalDocumentType>Identity document</personalDocumentType>
// 			<personalDocumentNumber>AAB123456</personalDocumentNumber>
// 			<birthDate>1984-01-18</birthDate>
// 			<birthPlace>Warsaw</birthPlace>
// 			<kind>OWNER</kind>
// 		</owner>
// 	</owners>
// 	<accounts>
// 		<account>
// 			<name>eKONTO</name>
// 			<iban>PL32114020040000320250132522</iban>
// 			<currencyBalance>1467.39</currencyBalance>
// 			<currencyFundsAvailable>1407.39</currencyFundsAvailable>
// 			<currencyName>PLN</currencyName>
// 			<owner>Jan Kowalski, Adam Nowak</owner> <!-- deprecated -->
// 			<activeSinceAtLeast>2000-02-28</activeSinceAtLeast>
// 			<owners>
// 				<owner>
// 				<name>Jan Kowalski</name>
// 				<address>Kwiatowa 8/97 Warszawa  PL</address>
// 				<kind>OWNER</kind>
// 				</owner>
// 			</owners>
// 			<owners>
// 				<owner>
// 				<name>Adam Nowak</name>
// 				<address>Robocza 1 M.14 Warszawa PL</address>
// 				<kind>CO-OWNER</kind>
// 				</owner>
// 			</owners>
// 		</account>
// 	</accounts>
// 	<moneyTransactions>
// 		<moneyTransaction>
// 		<transactionOn>2012-06-28</transactionOn>
// 		<bookedOn>2012-06-30</bookedOn>
// 		<currencyAmount>20.00</currencyAmount>
// 		<currencyBalance>230.00</currencyBalance>
// 		<title>Return for beer in a pub</title>
// 		<party>Jan Kowalski</party>
// 		<partyIban>PL68249000050000400075212326</partyIban>
// 		<kind>EXTERNAL INCOMING TRANSFER</kind>
// 		<status>DONE</status>
// 		<labels>
// 			<label>internal</label>
// 		</labels>
// 		</moneyTransaction>
// 		<moneyTransaction>
// 			<transactionOn>2017-08-10</transactionOn>
// 			<bookedOn>2017-08-11</bookedOn>
// 			<currencyAmount>-20.00</currencyAmount>
// 			<currencyBalance>1100.99</currencyBalance>
// 			<title>Zwrot za Sylwestra</title>
// 			<party>Jan Kowalski Dunikowskiego 23B, 01-200 Warszawa</party>
// 			<partyIban>PL83130000002076700146310001</partyIban>
// 			<kind>PRZELEW ZEWNĘTRZNY</kind>
// 			<labels>
// 				<label>internal</label>
// 			</labels>
// 		</moneyTransaction>
// 	</moneyTransactions>
// 	<creditCards>
// 		<creditCard>
// 		<name>Visa</name>
// 		<cardId>1234 1234 1234 1234</cardId>
// 		<iban>PL32114020040000320250132522</iban>
// 		<number>1234 1234 1234 1234</number>
// 		<currencyBalance>-2467.39</currencyBalance>
// 		<currencyName>PLN</currencyName>
// 		<owners>
// 			<owner>
// 				<name>Jan Kowalski</name>
// 				<address>Kwiatowa 8/97 Warszawa  PL</address>
// 				<kind>OWNER</kind>
// 			</owner>
// 			<owner>
// 				<name>Adam Nowak</name>
// 				<address>Robocza 1 M.14 Warszawa PL</address>
// 			</owner>
// 		</owners>
// 		<limit>5000.00</limit>
// 		<interest>20</interest>
// 		<dueDate>2013-10-26</dueDate>
// 		</creditCard>
// 	</creditCards>
// </target>`

// func TestKontomatikTestSuite(t *testing.T) {
// 	suite.Run(t, new(KontomatikTestSuite))
// }

// func (s *KontomatikTestSuite) SetupSuite() {
// 	err := dockerPostgresTest(s)
// 	if err != nil {
// 		log.Fatal().Msgf("failed to connect postgres: %v", err)
// 	}
// 	postgresPort := s.resource.GetPort("5432/tcp")

// 	cfg := &config.Config{
// 		PostgresHost:     "localhost",
// 		PostgresPort:     postgresPort,
// 		PostgresUsername: "postgres",
// 		PostgresPassword: "secret",
// 		PostgresDBName:   "test",
// 		PostgresSSLMode:  "disable",
// 	}
// 	s.cfg = cfg
// 	err = runMigrationsFromBinData(cfg)
// 	if err != nil {
// 		log.Fatal().Msg(err.Error())
// 	}

// 	s.repo = repoPostgres.NewKontomatikRepository(cfg)
// 	if err := s.repo.Connect(); err != nil {
// 		log.Fatal().Msgf("DB connection was not established: %+v", err)
// 		return
// 	}
// 	s.commands = getCommandsMockObjects()
// 	commands := make(map[string]command.Command)
// 	for k, v := range s.commands {
// 		commands[k] = v
// 	}
// 	importState := state.NewImportStateMachine(s.repo, s.cfg, commands)
// 	srv := NewKontomatikService(s.repo, importState, s.cfg, nil)
// 	s.srv = srv
// }

// func (s *KontomatikTestSuite) TearDownSuite() {
// 	defer func() {
// 		s.srv.rp.Close()
// 		s.db.Close()
// 		if err := s.pool.Purge(s.resource); err != nil {
// 			log.Fatal().Msgf("failed to remove postgresql container: %v", err)
// 		}
// 	}()
// }

// func (s *KontomatikTestSuite) TearDownTest() {
// 	_, err := s.db.Exec("DELETE FROM kontomatik_requests;")
// 	if err != nil {
// 		log.Fatal().Msgf("error delete from tables: %v", err)
// 	}
// 	s.commands = getCommandsMockObjects()
// 	commands := make(map[string]command.Command)
// 	for k, v := range s.commands {
// 		commands[k] = v
// 	}
// 	importState := state.NewImportStateMachine(s.repo, s.cfg, commands)
// 	srv := NewKontomatikService(s.repo, importState, s.cfg, nil)
// 	s.srv = srv
// }

// func dockerPostgresTest(s *KontomatikTestSuite) error {
// 	var db *sqlx.DB
// 	var err error
// 	pool, err := dockertest.NewPool("")
// 	if err != nil {
// 		log.Fatal().Msgf("Could not connect to docker: %s", err)
// 		return err
// 	}
// 	resource, err := pool.Run("postgres", "11.0", []string{"POSTGRES_PASSWORD=secret", "POSTGRES_DB=test"})
// 	if err != nil {
// 		log.Fatal().Msgf("Could not start resource: %s", err)
// 		return err
// 	}

// 	if err = pool.Retry(func() error {
// 		var internalErr error
// 		dataSourceName := fmt.Sprintf("postgres://postgres:secret@localhost:%s/%s?sslmode=disable", resource.GetPort("5432/tcp"), "test")
// 		db, internalErr = sqlx.Open("postgres", dataSourceName)
// 		if err != nil {
// 			return internalErr
// 		}
// 		return db.Ping()
// 	}); err != nil {
// 		log.Fatal().Msgf("Could not connect to docker: %s", err)
// 		return err
// 	}

// 	s.pool = pool
// 	s.resource = resource
// 	s.db = db

// 	return nil
// }

// func runMigrationsFromBinData(cfg *config.Config) error {
// 	binInstance, err := bindata.WithInstance(bindata.Resource(schema.AssetNames(), schema.Asset))
// 	if err != nil {
// 		log.Error().Err(err).Msg("failed to init db instance")
// 		return err
// 	}

// 	cfg.MigrationVersion = 9

// 	var conn *sqlx.DB

// 	dsn, err := cfg.GetDSN("postgres")
// 	if err != nil {
// 		return err
// 	}

// 	for tries := 0; tries <= 20; tries++ {
// 		conn, err = sqlx.Connect("postgres", dsn)
// 		if err != nil {
// 			log.Error().Err(err).Msg("failed to connect to DB. Retying in 2 seconds")
// 			time.Sleep(2000 * time.Millisecond)
// 			if tries == 20 {
// 				return err
// 			}
// 			continue
// 		}
// 		break
// 	}
// 	defer func() {
// 		if err = conn.Close(); err != nil {
// 			log.Error().Err(err).Msg("failed to close postgres connection")
// 		}
// 	}()

// 	targetInstance, err := postgres.WithInstance(conn.DB, new(postgres.Config))
// 	if err != nil {
// 		return err
// 	}
// 	m, err := migrate.NewWithInstance("go-bindata", binInstance, "postgres", targetInstance)
// 	if err != nil {
// 		return err
// 	}
// 	err = m.Migrate(cfg.MigrationVersion)
// 	if err != nil && err == migrate.ErrNoChange {
// 		log.Debug().Msg("No new migrations found")
// 		return nil
// 	} else if err != nil {
// 		log.Error().Err(err)
// 		return err
// 	} else {
// 		log.Debug().Msgf("Migrations to revision %d run.", cfg.MigrationVersion)
// 		return nil
// 	}
// }

// func getCommandsMockObjects() map[string]*CommandMockObject {
// 	commands := map[string]*CommandMockObject{
// 		consts.ImportOwnerRequestStatus:                 new(CommandMockObject),
// 		consts.ImportAccountRequestStatus:               new(CommandMockObject),
// 		consts.ImportAccountTransactionRequestStatus:    new(CommandMockObject),
// 		consts.ImportCreditCardRequestStatus:            new(CommandMockObject),
// 		consts.ImportCreditCardTransactionRequestStatus: new(CommandMockObject),
// 	}
// 	resp := &kontomatik.ReplyData{}
// 	resp.Command.State = kontomatik.StateSuccessful
// 	for i := range commands {
// 		commands[i].On("StartImportData", mock.Anything).Return("", nil)
// 		commands[i].On("GetDataByCommandID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
// 			Return(resp, nil)
// 		commands[i].On("ProcessData", mock.Anything, mock.Anything, mock.Anything).Return(nil)
// 	}
// 	return commands
// }

// func (s *KontomatikTestSuite) TestCreateNewKontomatikRequestRecord() {
// 	data := restapi.InitializeDataJSONRequestBody{
// 		ApplicationId:      gouuid.NewV1().String(),
// 		PartyId:            gouuid.NewV1().String(),
// 		SessionId:          gouuid.NewV1().String(),
// 		SessionIdSignature: gouuid.NewV1().String(),
// 	}

// 	res, err := s.srv.CreateKontomatikRequest(&data)
// 	if err != nil {
// 		s.FailNow("could not get data from service: ", err)
// 	}
// 	s.NotEmpty(res.Data)
// 	s.Equal(consts.NewRequestStatus, converto.StringValue(res.Data.Status))
// 	s.NotEmpty(res.Data.KontomatikRequestId)
// }

// func (s *KontomatikTestSuite) TestReturnExistsKontomatikRequestRecord() {
// 	data := restapi.InitializeDataJSONRequestBody{
// 		ApplicationId:      gouuid.NewV1().String(),
// 		PartyId:            gouuid.NewV1().String(),
// 		SessionId:          gouuid.NewV1().String(),
// 		SessionIdSignature: gouuid.NewV1().String(),
// 	}

// 	status := consts.ImportCreditCardRequestStatus
// 	s.insertKontomatikRequestTestData(data.ApplicationId, data.PartyId,
// 		data.SessionId, data.SessionIdSignature, status)

// 	res, err := s.srv.CreateKontomatikRequest(&data)

// 	s.NotEmpty(err)
// 	s.Nil(res)
// 	s.True(cerrors.IsExists(err))
// }

// func (s *KontomatikTestSuite) TestGetKontomatikRequestDataByID() {
// 	status := consts.NewRequestStatus
// 	id := s.insertKontomatikRequestTestData(gouuid.NewV1().String(), gouuid.NewV1().String(),
// 		gouuid.NewV1().String(), gouuid.NewV1().String(), status)
// 	s.insertAccountTestData(id)
// 	s.insertCreditCardTestData(id)

// 	res, err := s.srv.GetKontomatikRequest(id)
// 	if err != nil {
// 		s.FailNow("could not get data from service: ", err)
// 	}
// 	s.NotEmpty(res.Data)
// 	s.Equal(consts.NewRequestStatus, converto.StringValue(res.Data.Status))
// 	s.NotEmpty(res.Data.KontomatikRequestId)
// }

// func (s *KontomatikTestSuite) TestGetKontomatikRequestDataBySearchParams() {
// 	appID, partyID, status := gouuid.NewV1().String(), gouuid.NewV1().String(), consts.ImportFinishRequestStatus
// 	id := s.insertKontomatikRequestTestData(appID, partyID, gouuid.NewV1().String(), gouuid.NewV1().String(), status)
// 	s.insertKontomatikRequestTestData(appID, partyID, gouuid.NewV1().String(),
// 		gouuid.NewV1().String(), consts.ImportFailedRequestStatus)
// 	s.insertKontomatikRequestTestData(gouuid.NewV1().String(), partyID, gouuid.NewV1().String(),
// 		gouuid.NewV1().String(), status)
// 	s.insertKontomatikRequestTestData(appID, gouuid.NewV1().String(), gouuid.NewV1().String(),
// 		gouuid.NewV1().String(), status)
// 	s.insertKontomatikRequestTestData(gouuid.NewV1().String(), partyID, gouuid.NewV1().String(),
// 		gouuid.NewV1().String(), consts.ImportFailedRequestStatus)
// 	s.insertKontomatikRequestTestData(appID, gouuid.NewV1().String(), gouuid.NewV1().String(),
// 		gouuid.NewV1().String(), consts.ImportFailedRequestStatus)
// 	s.insertKontomatikRequestTestData(gouuid.NewV1().String(), gouuid.NewV1().String(), gouuid.NewV1().String(),
// 		gouuid.NewV1().String(), status)
// 	s.insertKontomatikRequestTestData(gouuid.NewV1().String(), gouuid.NewV1().String(), gouuid.NewV1().String(),
// 		gouuid.NewV1().String(), consts.ImportFailedRequestStatus)

// 	res, err := s.srv.GetKontomatikRequestByParams(restapi.GetKontomatikRequestByFiltersParams{
// 		ApplicationId: converto.StringPointer(appID),
// 		PartyId:       converto.StringPointer(partyID),
// 		Status:        converto.StringPointer(status),
// 	})
// 	if err != nil {
// 		s.FailNow("could not get data from service: ", err)
// 	}
// 	s.NotEmpty(res.Data)
// 	s.Equal(1, len(res.Data))
// 	req := res.Data[0]
// 	s.Equal(status, converto.StringValue(req.Status))
// 	s.Equal(appID, converto.StringValue(req.ApplicationId))
// 	s.Equal(partyID, converto.StringValue(req.PartyId))
// 	s.Equal(id.String(), converto.StringValue(req.KontomatikRequestId))
// }

// func (s *KontomatikTestSuite) TestGetPartyIncomeProofDataWithCorrectPartyID() {
// 	appID, partyID := gouuid.NewV1(), gouuid.NewV1()
// 	id := s.insertKontomatikRequestTestData(appID.String(), partyID.String(), gouuid.NewV1().String(),
// 		gouuid.NewV1().String(), consts.ImportFinishRequestStatus)
// 	s.insertKontomatikRequestTestData(gouuid.NewV1().String(), partyID.String(), gouuid.NewV1().String(),
// 		gouuid.NewV1().String(), consts.ImportFailedRequestStatus)
// 	s.insertKontomatikRequestTestData(appID.String(), gouuid.NewV1().String(), gouuid.NewV1().String(),
// 		gouuid.NewV1().String(), consts.ImportFailedRequestStatus)
// 	s.updateRequestOverallData(id)

// 	res, err := s.srv.GetPartyIncomeProofData(partyID, appID.String())
// 	if err != nil {
// 		s.FailNow("could not get data from service: ", err)
// 	}
// 	s.NotEmpty(res.Data)
// 	s.Equal(partyID.String(), converto.StringValue(res.Data.PartyId))
// 	s.Equal(appID.String(), converto.StringValue(res.Data.ApplicationId))
// 	s.NotEmpty(res.Data.OverallReports)
// }

// func (s *KontomatikTestSuite) TestGetPartyIncomeProofDataWithUncorrectPartyID() {
// 	partyID := gouuid.NewV1()
// 	id := s.insertKontomatikRequestTestData(gouuid.NewV1().String(), partyID.String(), gouuid.NewV1().String(),
// 		gouuid.NewV1().String(), consts.ImportFinishRequestStatus)
// 	s.insertKontomatikRequestTestData(gouuid.NewV1().String(), partyID.String(), gouuid.NewV1().String(),
// 		gouuid.NewV1().String(), consts.ImportFailedRequestStatus)
// 	s.updateRequestOverallData(id)

// 	res, err := s.srv.GetPartyIncomeProofData(gouuid.NewV1(), gouuid.NewV1().String())
// 	if res != nil {
// 		s.FailNow("could not get data from service: ", err)
// 	}
// 	s.NotEmpty(err)
// 	s.True(cerrors.IsNotExist(err))
// }

// func (s *KontomatikTestSuite) TestStartLoadDataFromDifferentStatuses() {
// 	s.assertCallsCountByStatus(consts.NewRequestStatus, map[string]int{
// 		consts.ImportOwnerRequestStatus:                 1,
// 		consts.ImportAccountRequestStatus:               1,
// 		consts.ImportAccountTransactionRequestStatus:    1,
// 		consts.ImportCreditCardRequestStatus:            1,
// 		consts.ImportCreditCardTransactionRequestStatus: 1,
// 	})
// }

// func (s *KontomatikTestSuite) TestKontomatikGetScoreByPartyID() {
// 	id, partyID := s.insertPartyScoreTestData()

// 	res, err := s.srv.GetScoreByPartyID(partyID, partyID.String())
// 	if err != nil {
// 		s.FailNow("could not get data score from service: ", err)
// 	}
// 	s.Equal(id.String(), converto.StringValue(res.Data.Id))
// 	s.Equal(partyID.String(), converto.StringValue(res.Data.PartyId))
// }

// func (s *KontomatikTestSuite) TestKontomatikGetScoreByUncorrectPartyID() {
// 	s.insertPartyScoreTestData()

// 	res, err := s.srv.GetScoreByPartyID(gouuid.NewV4(), gouuid.NewV4().String())
// 	if res != nil {
// 		s.FailNow("not correct data for GetScoreByPartyID: ", err)
// 	}
// 	s.NotEmpty(err)
// 	s.True(cerrors.IsNotExist(err))
// }

// func (s *KontomatikTestSuite) assertCallsCountByStatus(status string, callsCount map[string]int) {
// 	id := s.insertKontomatikRequestTestData(gouuid.NewV1().String(), gouuid.NewV1().String(),
// 		gouuid.NewV1().String(), gouuid.NewV1().String(), status)
// 	s.insertAccountTestData(id)
// 	s.insertCreditCardTestData(id)

// 	kr, err := s.repo.GetRequestByID(id)
// 	if err != nil {
// 		s.FailNow("failed to get kontomatik request by id", err)
// 	}
// 	s.srv.startLoadData(kr)

// 	for i := range s.commands {
// 		s.commands[i].AssertNumberOfCalls(s.T(), "StartImportData", callsCount[i])
// 		s.commands[i].AssertNumberOfCalls(s.T(), "GetDataByCommandID", callsCount[i])
// 		s.commands[i].AssertNumberOfCalls(s.T(), "ProcessData", callsCount[i])
// 	}

// 	err = s.db.Get(&status, "SELECT status FROM kontomatik_requests WHERE id = $1", id)
// 	if err != nil {
// 		s.FailNow("failed to get created consent with expected parameters", err)
// 	}
// 	jobCount := 0
// 	err = s.db.Get(&jobCount, "SELECT count(1) FROM kontomatik_request_jobs WHERE kontomatik_request_id = $1 and status = $2",
// 		id, consts.JobStatusSuccess)
// 	if err != nil {
// 		s.FailNow("failed to get created consent with expected parameters", err)
// 	}
// 	s.Equal(consts.ImportFinishRequestStatus, status)
// 	s.Equal(3, jobCount)
// }

// func (s *KontomatikTestSuite) insertKontomatikRequestTestData(appID, partyID, sessionID, signature, status string) gouuid.UUID {
// 	id := gouuid.NewV1()
// 	_, err := s.db.Exec(`INSERT INTO public.kontomatik_requests (id, created_at, updated_at, application_id, party_id,
// 		session_id, session_id_signature, status, overall_data)
// 		VALUES ($1, '2020-11-10 14:44:32.943587', '2020-11-10 14:44:32.943587', $2, $3, $4, $5, $6, null);`,
// 		id, appID, partyID, sessionID, signature, status)
// 	if err != nil {
// 		s.FailNow("could not insert test client documents ", err)
// 	}
// 	return id
// }

// func (s *KontomatikTestSuite) insertAccountTestData(reqID gouuid.UUID) {
// 	id := gouuid.NewV1()
// 	_, err := s.db.Exec(`INSERT INTO public.accounts (id, created_at, kontomatik_request_id, iban, card_id,
// 		currency_balance, funds_available)
// 		VALUES ($1, '2020-11-10 15:30:10.961965',
// 		$2, 'PL03800400022007001839630001', null, null, null);`,
// 		id, reqID)
// 	if err != nil {
// 		s.FailNow("could not insert test client documents ", err)
// 	}
// }

// func (s *KontomatikTestSuite) insertCreditCardTestData(reqID gouuid.UUID) {
// 	id := gouuid.NewV1()
// 	_, err := s.db.Exec(`INSERT INTO public.accounts (id, created_at, kontomatik_request_id, iban, card_id,
// 		currency_balance, funds_available)
// 		VALUES ($1, '2020-11-10 15:30:34.913530',
// 		$2, 'PL19025714369262645919804211', '**************13',
// 		-78.98, null);`,
// 		id, reqID)
// 	if err != nil {
// 		s.FailNow("could not insert test client documents ", err)
// 	}
// }

// func (s *KontomatikTestSuite) updateRequestOverallData(reqID gouuid.UUID) {
// 	data := testOverallData
// 	overallData := map[string]string{
// 		resultXMLDataKey: data,
// 	}
// 	err := s.repo.MergeRequestOverallData(reqID, overallData)
// 	if err != nil {
// 		s.FailNow("failed write response to db for kontomatik request id = %+v with error: %+v", reqID, err)
// 	}
// }

// // Insert party score and return id and party_id.
// // appliaction_id in score will be same as party_id
// func (s *KontomatikTestSuite) insertPartyScoreTestData() (id, partyID gouuid.UUID) {
// 	id, partyID = gouuid.NewV4(), gouuid.NewV4()
// 	_, err := s.db.Exec(`INSERT INTO scores (id,created_at,updated_at,party_id,score,score_percentile,
// 		score_tier,overall_data, application_id)
// 		VALUES ($1,
// 		'2020-12-29 11:58:23.345919','2020-12-29 11:58:23.345919',$2,
// 		0.21170372,0.07035766,'F',null, $3)`, id, partyID, partyID.String())
// 	if err != nil {
// 		s.FailNow("could not insert test client documents ", err)
// 	}
// 	return id, partyID
// }
