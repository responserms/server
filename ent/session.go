// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/responserms/server/ent/session"
	"github.com/responserms/server/ent/token"
	"github.com/responserms/server/ent/user"
)

// Session is the model entity for the Session schema.
type Session struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// IPAddress holds the value of the "ip_address" field.
	IPAddress string `json:"ip_address,omitempty"`
	// BrowserName holds the value of the "browser_name" field.
	BrowserName string `json:"browser_name,omitempty"`
	// BrowserVersion holds the value of the "browser_version" field.
	BrowserVersion string `json:"browser_version,omitempty"`
	// DeviceOs holds the value of the "device_os" field.
	DeviceOs string `json:"device_os,omitempty"`
	// DeviceType holds the value of the "device_type" field.
	DeviceType string `json:"device_type,omitempty"`
	// TerminatedAt holds the value of the "terminated_at" field.
	TerminatedAt *time.Time `json:"terminated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SessionQuery when eager-loading is set.
	Edges         SessionEdges `json:"edges"`
	token_session *int
	user_sessions *int
}

// SessionEdges holds the relations/edges for other nodes in the graph.
type SessionEdges struct {
	// Token holds the value of the token edge.
	Token *Token
	// User holds the value of the user edge.
	User *User
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TokenOrErr returns the Token value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SessionEdges) TokenOrErr() (*Token, error) {
	if e.loadedTypes[0] {
		if e.Token == nil {
			// The edge token was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: token.Label}
		}
		return e.Token, nil
	}
	return nil, &NotLoadedError{edge: "token"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SessionEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Session) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case session.FieldID:
			values[i] = &sql.NullInt64{}
		case session.FieldIPAddress, session.FieldBrowserName, session.FieldBrowserVersion, session.FieldDeviceOs, session.FieldDeviceType:
			values[i] = &sql.NullString{}
		case session.FieldCreateTime, session.FieldUpdateTime, session.FieldTerminatedAt:
			values[i] = &sql.NullTime{}
		case session.ForeignKeys[0]: // token_session
			values[i] = &sql.NullInt64{}
		case session.ForeignKeys[1]: // user_sessions
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Session", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Session fields.
func (s *Session) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case session.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case session.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				s.CreateTime = value.Time
			}
		case session.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				s.UpdateTime = value.Time
			}
		case session.FieldIPAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ip_address", values[i])
			} else if value.Valid {
				s.IPAddress = value.String
			}
		case session.FieldBrowserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field browser_name", values[i])
			} else if value.Valid {
				s.BrowserName = value.String
			}
		case session.FieldBrowserVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field browser_version", values[i])
			} else if value.Valid {
				s.BrowserVersion = value.String
			}
		case session.FieldDeviceOs:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field device_os", values[i])
			} else if value.Valid {
				s.DeviceOs = value.String
			}
		case session.FieldDeviceType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field device_type", values[i])
			} else if value.Valid {
				s.DeviceType = value.String
			}
		case session.FieldTerminatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field terminated_at", values[i])
			} else if value.Valid {
				s.TerminatedAt = new(time.Time)
				*s.TerminatedAt = value.Time
			}
		case session.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field token_session", value)
			} else if value.Valid {
				s.token_session = new(int)
				*s.token_session = int(value.Int64)
			}
		case session.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_sessions", value)
			} else if value.Valid {
				s.user_sessions = new(int)
				*s.user_sessions = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryToken queries the token edge of the Session.
func (s *Session) QueryToken() *TokenQuery {
	return (&SessionClient{config: s.config}).QueryToken(s)
}

// QueryUser queries the user edge of the Session.
func (s *Session) QueryUser() *UserQuery {
	return (&SessionClient{config: s.config}).QueryUser(s)
}

// Update returns a builder for updating this Session.
// Note that, you need to call Session.Unwrap() before calling this method, if this Session
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Session) Update() *SessionUpdateOne {
	return (&SessionClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Session) Unwrap() *Session {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Session is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Session) String() string {
	var builder strings.Builder
	builder.WriteString("Session(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(s.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(s.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ip_address=")
	builder.WriteString(s.IPAddress)
	builder.WriteString(", browser_name=")
	builder.WriteString(s.BrowserName)
	builder.WriteString(", browser_version=")
	builder.WriteString(s.BrowserVersion)
	builder.WriteString(", device_os=")
	builder.WriteString(s.DeviceOs)
	builder.WriteString(", device_type=")
	builder.WriteString(s.DeviceType)
	if v := s.TerminatedAt; v != nil {
		builder.WriteString(", terminated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Sessions is a parsable slice of Session.
type Sessions []*Session

func (s Sessions) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
