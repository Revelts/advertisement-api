package advertisement

const (
	CreateAdvertisement                     = `INSERT INTO public.ads (advertisement_name, advertisement_category, advertisement_baseprice, advertisement_created_by) VALUES($1, $2, $3, $4) RETURNING advertisement_id`
	ViewAllAdvertisements                   = `SELECT advertisement_id, advertisement_owner, advertisement_name, advertisement_category, advertisement_baseprice, advertisement_created_by, TO_CHAR(advertisement_created_at,'YYYY-MM-DD HH24:MI:SS'), TO_CHAR(advertisement_updated_at,'YYYY-MM-DD HH24:MI:SS')  FROM public.ads`
	FailToCreateAdvertisementNotEnoughMoney = `You don't have enough budget to place an advertisement'`
	FailToBuyAdvertisementNotEnoughMoney    = `You don't have enough budget to buy this advertisement'`
	GetAdvertisementInfo                    = `SELECT advertisement_id, advertisement_owner, advertisement_name, advertisement_category, advertisement_baseprice, advertisement_created_by FROM public.ads WHERE advertisement_id = $1`
	BuyAdvertisement                        = `UPDATE public.ads SET advertisement_owner = $1, advertisement_updated_at = now() WHERE advertisement_id = $2`
)
